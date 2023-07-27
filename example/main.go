package main

import (
	"fmt"
	"os"
	"time"

	"github.com/verystar/otp/totp"
)

func main() {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Example.com",
		AccountName: "alice@example.com",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Issuer:       %s\n", key.Issuer())
	fmt.Printf("Account Name: %s\n", key.AccountName())
	fmt.Printf("Secret:       %s\n", key.Secret())
	fmt.Printf("URL:       %s\n", key.String())
	fmt.Println("")

	passcode, err := totp.GenerateCode(key.Secret(), time.Now())

	if err != nil {
		panic(err)
	}

	fmt.Println("Passcode:", passcode)

	// Now Validate that the user's successfully added the passcode.
	valid := totp.Validate(passcode, key.Secret())
	if valid {
		println("Success passcode!")
		os.Exit(0)
	} else {
		println("Invalid passcode!")
		os.Exit(1)
	}
}
