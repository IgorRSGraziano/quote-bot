package whatsapp

import (
	"context"
	"fmt"

	"math/rand"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
)

func eventHandler(client *whatsmeow.Client) func(interface{}) {
	return func(evt interface{}) {
		switch v := evt.(type) {
		case *events.Message:
			handleMessage(client, v)
		default:
			fmt.Println("Received a message!", v)
		}

	}
}

func isFromEnabledGroup(msg *events.Message) bool {
	//TODO: Colocar em um env
	return msg.Info.MessageSource.Chat.User == "120363346101975622"
}

func isActivateCommand(msg *events.Message) bool {
	return msg.Message.GetConversation() == "!nathan"
}

func handleMessage(client *whatsmeow.Client, msg *events.Message) {
	if !isFromEnabledGroup(msg) || !isActivateCommand(msg) {
		return
	}

	quotes := []string{
		"Frase 1",
		"Frase 2",
		"Frase 3",
	}

	randomQuote := &quotes[rand.Intn(len(quotes))]

	ctx := context.Background()

	//TODO: Colocar quotedMessage
	r, err := client.SendMessage(ctx, msg.Info.Chat, &waE2E.Message{
		Conversation: randomQuote,
	})

	fmt.Println(r)

	if err != nil {
		client.Log.Errorf("Error sending message: %v", err)
	}
}
