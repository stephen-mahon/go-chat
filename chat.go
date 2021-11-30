package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ably/ably-go/ably"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error loading .env file:", err)
	}

	var username string

	// If no username specified, ask for one
	if len(os.Args) < 2 {
		fmt.Println("Type your username")
		reader := bufio.NewReader(os.Stdin)
		username, _ = reader.ReadString('\n')
		username = strings.Replace(username, "\n", "", -1)
	} else {
		username = os.Args[1]
	}

	// Connect to Ably using the API key and ClientID specified above
	client, err := ably.NewRealtime(
		ably.WithKey(os.Getenv("ABLY_KEY")),
		// ably.WithEchoMessages(true), // Uncomment to stop messages you send from being sent back
		ably.WithClientID(username))
	if err != nil {
		panic(err)
	}

	fmt.Println("You can now send messages!")
	_ = client

}
