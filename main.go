package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()
}

func mustToken() string {
	token := flag.String("token-tg", "", "token to access to the telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token must be specified")
	}

	return *token
}
