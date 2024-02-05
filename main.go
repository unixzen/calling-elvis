package main

import (
	"flag"
	"log"

	"github.com/unixzen/calling-elvis/clients/telegram"
)

const tgHost = "api.telegram.org"

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for telegram bot")
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
}

func main() {
	tgClient := telegram.New(tgHost, mustToken())

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)

}
