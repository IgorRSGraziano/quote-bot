package main

import (
	"fmt"
	"quote-bot/whatsapp"

	_ "quote-bot/config/autoload"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow/types/events"
)

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
	}
}

func main() {
	whatsapp.Start()
}
