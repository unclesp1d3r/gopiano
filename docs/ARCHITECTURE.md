# gopiano Architecture

This document describes the architecture, design patterns, and internal workings of the gopiano library.

## System Overview

```mermaid
graph TB
    subgraph "Application Layer"
        App[Your Application]
    end

    subgraph "gopiano Library"
        Client[Client]
        Auth[auth.go]
        User[user.go]
        Station[station.go]
        Misc[misc.go]
        Core[gopiano.go]
    end

    subgraph "Data Layer"
        Requests[requests package]
        Responses[responses package]
    end

    subgraph "External"
        Pandora[Pandora API]
    end

    App --> Client
    Client --> Auth
    Client --> User
    Client --> Station
    Client --> Misc
    Auth --> Core
    User --> Core
    Station --> Core
    Misc --> Core
    Core --> Requests
    Core --> Responses
    Core -->|HTTPS/HTTP| Pandora
    Pandora -->|JSON| Core
```

## Package Structure

```
github.com/unclesp1d3r/gopiano/
├── gopiano.go          # Core client, encryption, HTTP calls
├── auth.go             # Authentication methods
├── user.go             # User-related methods
├── station.go          # Station management methods
├── misc.go             # Music search, bookmarks, track info
├── requests/
│   └── requests.go     # Request data structures for JSON marshaling
└── responses/
    └── responses.go    # Response data structures for JSON unmarshaling
```

## Component Diagram

```mermaid
classDiagram
    class Client {
        -description ClientDescription
        -http *http.Client
        -encrypter *blowfish.Cipher
        -decrypter *blowfish.Cipher
        -timeOffset time.Duration
        -partnerAuthToken string
        -partnerID string
        -userAuthToken string
        -userID string
        +NewClient(d) *Client
        +PandoraCall(ctx, protocol, method, body, data) error
        +BlowfishCall(ctx, protocol, method, body, data) error
        +GetSyncTime() int
        +AuthPartnerLogin(ctx) *AuthPartnerLogin
        +AuthUserLogin(ctx, user, pass) *AuthUserLogin
        +UserGetStationList(ctx, includeArt) *UserGetStationList
        +StationGetPlaylist(ctx, token) *StationGetPlaylist
    }

    class ClientDescription {
        +DeviceModel string
        +Username string
        +Password string
        +BaseURL string
        +EncryptKey string
        +DecryptKey string
        +Version string
    }

    class PandoraError {
        +Stat string
        +Code int
        +Message string
        +Error() string
    }

    Client --> ClientDescription : uses
    Client ..> PandoraError : returns
```

## Authentication Flow

```mermaid
sequenceDiagram
    participant App as Application
    participant Client as gopiano.Client
    participant API as Pandora API

    App->>Client: NewClient(AndroidClient)
    Client-->>App: *Client

    Note over App,API: Step 1: Partner Authentication (Required First)
    App->>Client: AuthPartnerLogin(ctx)
    Client->>API: POST auth.partnerLogin
    API-->>Client: partnerAuthToken, partnerID, syncTime
    Client->>Client: Store tokens, calculate timeOffset
    Client-->>App: *AuthPartnerLogin

    Note over App,API: Step 2: User Authentication
    App->>Client: AuthUserLogin(ctx, email, password)
    Client->>Client: Validate partnerAuthToken exists
    Client->>Client: Encrypt request body (Blowfish)
    Client->>API: POST auth.userLogin (encrypted)
    API-->>Client: userAuthToken, userID
    Client->>Client: Store user tokens
    Client-->>App: *AuthUserLogin

    Note over App,API: Step 3: Authenticated API Calls
    App->>Client: UserGetStationList(ctx, false)
    Client->>Client: Validate userAuthToken exists
    Client->>Client: Encrypt request body
    Client->>API: POST user.getStationList (encrypted)
    API-->>Client: Station list JSON
    Client-->>App: *UserGetStationList
```

## Encryption Architecture

The Pandora API uses Blowfish encryption in ECB (Electronic Codebook) mode for request bodies.

```mermaid
flowchart LR
    subgraph "Request Encryption"
        A[JSON Request] --> B[Blowfish ECB Encrypt]
        B --> C[Hex Encode]
        C --> D[HTTP POST Body]
    end

    subgraph "Response Decryption"
        E[HTTP Response] --> F[JSON Parse]
        F --> G{Encrypted Field?}
        G -->|Yes| H[Hex Decode]
        H --> I[Blowfish ECB Decrypt]
        I --> J[Process Data]
        G -->|No| J
    end
```

### Encryption Details

- **Algorithm**: Blowfish cipher
- **Mode**: ECB (Electronic Codebook)
- **Key Size**: Variable (provided by ClientDescription)
- **Encoding**: Hex encoding for transport

```go
// Encryption process (simplified)
func (c *Client) encrypt(data string) string {
    chunks := []string{}
    for i := 0; i < len(data); i += 8 {
        var buf, crypt [8]byte
        copy(buf[:], data[i:])
        c.encrypter.Encrypt(crypt[:], buf[:])
        chunks = append(chunks, hex.EncodeToString(crypt[:]))
    }
    return strings.Join(chunks, "")
}
```

## Request/Response Flow

```mermaid
flowchart TB
    subgraph "API Method Call"
        A[Method Called] --> B{Auth Required?}
        B -->|Partner| C[validatePartnerAuthToken]
        B -->|User| D[validateUserAuthToken]
        C --> E{Token Exists?}
        D --> E
        E -->|No| F[Return Error]
        E -->|Yes| G[Create Request Struct]
    end

    subgraph "Request Processing"
        G --> H[json.Marshal]
        H --> I{Encryption Needed?}
        I -->|Yes| J[BlowfishCall]
        I -->|No| K[PandoraCall]
        J --> L[encrypt body]
        L --> K
    end

    subgraph "HTTP Layer"
        K --> M[Build URL with params]
        M --> N[http.NewRequestWithContext]
        N --> O[Set Headers]
        O --> P[http.Client.Do]
    end

    subgraph "Response Processing"
        P --> Q[Read Response Body]
        Q --> R[json.Unmarshal to error check]
        R --> S{stat == fail?}
        S -->|Yes| T[Return PandoraError]
        S -->|No| U[json.Unmarshal to response]
        U --> V[Return Response]
    end
```

## State Management

The `Client` struct maintains authentication state across API calls:

```mermaid
stateDiagram-v2
    [*] --> Uninitialized: NewClient()

    Uninitialized --> PartnerAuth: AuthPartnerLogin()
    note right of PartnerAuth
        Sets:
        - partnerAuthToken
        - partnerID
        - timeOffset
    end note

    PartnerAuth --> UserAuth: AuthUserLogin() or UserCreateUser()
    note right of UserAuth
        Sets:
        - userAuthToken
        - userID
    end note

    UserAuth --> UserAuth: API Calls
    note right of UserAuth
        Can call:
        - UserGetStationList
        - StationGetPlaylist
        - MusicSearch
        - etc.
    end note
```

## URL Construction

API calls construct URLs with query parameters:

```
https://tuner.pandora.com/services/json/?method={method}&partner_id={id}&user_id={id}&auth_token={token}
```

```mermaid
flowchart LR
    A[Base URL] --> B[Add method param]
    B --> C{partnerID set?}
    C -->|Yes| D[Add partner_id]
    C -->|No| E[Skip]
    D --> F{userID set?}
    E --> F
    F -->|Yes| G[Add user_id]
    F -->|No| H[Skip]
    G --> I{Auth token?}
    H --> I
    I -->|Partner only| J[Add partnerAuthToken]
    I -->|User| K[Add userAuthToken]
    J --> L[Final URL]
    K --> L
```

## Error Handling Strategy

```mermaid
flowchart TB
    A[API Call] --> B{Client-side validation}
    B -->|Missing auth| C[Return Go error]
    B -->|Valid| D[Make HTTP request]
    D --> E{HTTP error?}
    E -->|Yes| F[Return HTTP error]
    E -->|No| G[Parse JSON]
    G --> H{stat == fail?}
    H -->|Yes| I[Create PandoraError]
    I --> J{Code == 0?}
    J -->|Yes| K[Add guidance message]
    J -->|No| L[Use ErrorCodeMap]
    K --> M[Return PandoraError]
    L --> M
    H -->|No| N[Parse response struct]
    N --> O[Return success]
```

## Concurrency Considerations

```mermaid
flowchart TB
    subgraph "NOT Thread-Safe"
        A[Client Instance]
        A --> B[partnerAuthToken]
        A --> C[userAuthToken]
        A --> D[timeOffset]
        A --> E[partnerID]
        A --> F[userID]
    end

    subgraph "Safe Patterns"
        G[Option 1: Separate Clients]
        H[Goroutine 1] --> I[Client 1]
        J[Goroutine 2] --> K[Client 2]

        L[Option 2: External Mutex]
        M[sync.Mutex] --> N[Shared Client]
        O[Goroutine 1] --> M
        P[Goroutine 2] --> M
    end
```

## Dependencies

```mermaid
graph TD
    subgraph "gopiano"
        Main[gopiano]
    end

    subgraph "Standard Library"
        Context[context]
        JSON[encoding/json]
        HTTP[net/http]
        URL[net/url]
        Time[time]
        IO[io]
        Hex[encoding/hex]
        Fmt[fmt]
        Strings[strings]
    end

    subgraph "External"
        Blowfish[golang.org/x/crypto/blowfish]
    end

    Main --> Context
    Main --> JSON
    Main --> HTTP
    Main --> URL
    Main --> Time
    Main --> IO
    Main --> Hex
    Main --> Fmt
    Main --> Strings
    Main --> Blowfish
```

## Design Decisions

### 1. Thin Wrapper Approach

The library is intentionally a "thin wrapper" that mirrors the Pandora API exactly:

- Request/response structs match API JSON structure
- Method names match API method names
- No business logic abstraction

### 2. Context Support

All API methods accept `context.Context` as the first parameter:

- Enables request cancellation
- Supports timeouts
- Allows request tracing

### 3. Client-Side Validation

Authentication state is validated before making API calls:

- Prevents cryptic server errors
- Provides actionable error messages
- Fails fast on misuse

### 4. Blowfish Encryption

Uses deprecated `golang.org/x/crypto/blowfish` package:

- Required by Pandora's API protocol
- Not a security concern (API requirement, not design choice)
- Annotated with `//nolint:staticcheck`

### 5. Separation of Concerns

- `gopiano.go`: Core HTTP and encryption infrastructure
- `auth.go`: Authentication methods
- `user.go`: User management methods
- `station.go`: Station operations
- `misc.go`: Search, bookmarks, track info
- `requests/`: Data structures for outgoing requests
- `responses/`: Data structures for incoming responses
