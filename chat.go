package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ably/ably-go/ably"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	username := flag.String("u", "Steve", "username")
	flag.Parse()

	// Connect to Ably using the API key and ClientID specified above
	client, err := ably.NewRealtime(
		ably.WithKey(os.Getenv("ABLY_KEY")),
		// ably.WithEchoMessages(true), // Uncomment to stop messages you send from being sent back
		ably.WithClientID(*username))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s, you can now send messages!\n", client.Auth.ClientID())
	_ = client

}
