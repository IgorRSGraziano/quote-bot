package whatsapp

import (
	"fmt"

	"go.mau.fi/whatsmeow/types/events"
)

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
	}
}

func handleMessage(msg *events.Message) {
	fmt.Println("Received a message!", msg.Message.GetConversation())
}