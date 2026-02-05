// Package examples demonstrates proper usage of the gopiano library.
//
// This example shows the basic authentication flow for existing users.
// This is the minimal authentication flow that must be completed before
// calling any user-specific API methods.
//
// ⚠️ Disclaimer: This example is for educational and demonstration purposes only.
// Users must have valid Pandora account credentials and are responsible for
// ensuring they have legal rights to access the Pandora API and comply with
// Pandora's Terms of Service. The placeholder credentials in this example must
// be replaced with actual valid credentials.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/unclesp1d3r/gopiano"
)

func main() {
	// Create a context for API calls. You can add timeout or cancellation as needed.
	ctx := context.Background()

	// Step 1: Create a client using the Android client description.
	// This provides the necessary encryption keys and device model information
	// that Pandora's API expects.
	client, err := gopiano.NewClient(gopiano.AndroidClient)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Step 2: Establish partner session (REQUIRED FIRST).
	// This MUST be called before any other API methods. It:
	// - Authenticates with Pandora using partner credentials
	// - Obtains partnerAuthToken and partnerID
	// - Sets timeOffset for sync time calculations
	// These values are stored on the client and used automatically in subsequent calls.
	partnerResp, err := client.AuthPartnerLogin(ctx)
	if err != nil {
		log.Fatalf("Failed to authenticate partner: %v\n"+
			"Note: This may fail if you're not calling from a US IP address.", err)
	}
	fmt.Printf("Partner authenticated successfully.\n")
	fmt.Printf("  Partner ID: %s\n", partnerResp.Result.PartnerID)
	fmt.Printf("  Partner Auth Token: %s\n", partnerResp.Result.PartnerAuthToken)

	// Step 3: Authenticate as an existing user.
	// This requires valid Pandora account credentials (email and password).
	// After this call succeeds, the client will have userAuthToken and userID set,
	// allowing you to make authenticated user API calls.
	//
	// ⚠️ IMPORTANT: Replace these placeholder credentials with your actual valid
	// Pandora account credentials. You must have legal rights to use these credentials
	// and must comply with Pandora's Terms of Service.
	username := "user@example.com"
	password := "your-password"

	userResp, err := client.AuthUserLogin(ctx, username, password)
	if err != nil {
		log.Fatalf("Failed to authenticate user: %v\n"+
			"Common causes:\n"+
			"  - Invalid username or password\n"+
			"  - Missing AuthPartnerLogin() call (should be caught by validation)\n"+
			"  - Not calling from US IP address\n"+
			"  - Account may be locked or disabled", err)
	}

	// Step 4: Verify authentication succeeded.
	// The client now has all necessary tokens for making authenticated API calls.
	fmt.Printf("User authenticated successfully.\n")
	fmt.Printf("  Username: %s\n", userResp.Result.Username)
	fmt.Printf("  User ID: %s\n", userResp.Result.UserID)
	fmt.Printf("  Can Listen: %v\n", userResp.Result.CanListen)
	fmt.Printf("  Max Stations Allowed: %d\n", userResp.Result.MaxStationsAllowed)

	// The client is now fully authenticated and ready for user-specific API calls.
	// For example:
	//   stations, err := client.UserGetStationList(ctx, false)
	//   if err != nil {
	//       log.Fatal(err)
	//   }
	//   fmt.Printf("User has %d stations\n", len(stations.Result.Stations))
}
