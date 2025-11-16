// Package examples demonstrates proper usage of the gopiano library.
//
// This example shows how to create a new Pandora user account using UserCreateUser.
// Prerequisites: US IP address, valid parameters, and the legacy API must still support account creation.
package main

import (
	"fmt"
	"log"

	"github.com/unclesp1d3r/gopiano"
)

func main() {
	// Step 1: Create a client using the Android client description.
	// This provides the necessary encryption keys and device model information.
	client, err := gopiano.NewClient(gopiano.AndroidClient)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Step 2: Establish partner session (REQUIRED FIRST).
	// This MUST be called before any other API methods. It obtains the partner
	// authentication token that is required for subsequent calls, including UserCreateUser.
	partnerResp, err := client.AuthPartnerLogin()
	if err != nil {
		log.Fatalf("Failed to authenticate partner: %v\n"+
			"Note: This may fail if you're not calling from a US IP address.", err)
	}
	fmt.Printf("Partner authenticated. Partner ID: %s\n", partnerResp.Result.PartnerID)

	// Step 3: Create a new user account.
	// All parameters must meet Pandora's requirements:
	// - username: Must be a valid email address
	// - password: User's chosen password
	// - gender: Must be exactly "male" or "female"
	// - countryCode: Must be "US" (API restriction)
	// - zipCode: Must be a valid US ZIP code (5 digits)
	// - birthYear: Must meet minimum age requirements (typically 13+)
	// - emailOptin: Whether user opts in to marketing emails
	username := "user@example.com"
	password := "SecurePassword123"
	gender := "male"
	countryCode := "US"
	zipCode := 90210
	birthYear := 1990
	emailOptin := false

	userResp, err := client.UserCreateUser(
		username,
		password,
		gender,
		countryCode,
		zipCode,
		birthYear,
		emailOptin,
	)
	if err != nil {
		log.Fatalf("Failed to create user: %v\n"+
			"Common causes:\n"+
			"  - Missing AuthPartnerLogin() call (should be caught by validation)\n"+
			"  - Invalid parameters (email format, gender, zip code, birth year)\n"+
			"  - Not calling from US IP address\n"+
			"  - Rate limiting (too many requests)\n"+
			"  - Username already exists\n"+
			"  - API endpoint may be deprecated or restricted", err)
	}

	// Step 4: Success! The user account has been created.
	// The client now has userAuthToken and userID set, so you can make
	// authenticated user API calls immediately.
	fmt.Printf("User created successfully!\n")
	fmt.Printf("  Username: %s\n", userResp.Result.Username)
	fmt.Printf("  User ID: %s\n", userResp.Result.UserID)
	fmt.Printf("  Can Listen: %v\n", userResp.Result.CanListen)
	fmt.Printf("  Max Stations Allowed: %d\n", userResp.Result.MaxStationsAllowed)

	// The client is now fully authenticated and ready for user-specific API calls.
	// For example, you could now call:
	//   stations, err := client.UserGetStationList(false)
}
