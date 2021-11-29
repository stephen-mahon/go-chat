package main

import (
	"fmt"
	"os"

	"github.com/ably/ably-go/ably"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	username := os.Args[1]

	client, _ := ably.NewRealtime(
		ably.WithKey(os.Getenv("ABLY_KEY")),
		ably.WithClientID(username),
		ably.WithEchoMessages(false),
	)

	fmt.Println(client.Auth.ClientID())

}
