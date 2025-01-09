package whatsapp

import (
	"context"
	"fmt"
	"quote-bot/config"
	"quote-bot/utils"

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
		}

	}
}

func isFromEnabledGroup(msg *events.Message) bool {
	fmt.Println(msg.Info.MessageSource.Chat.User, config.Data.Group)
	return msg.Info.MessageSource.Chat.User == config.Data.Group
}

func isActivateCommand(msg *events.Message) bool {
	return msg.Message.GetConversation() == config.Data.Command
}

func handleMessage(client *whatsmeow.Client, event *events.Message) {
	if !isFromEnabledGroup(event) || !isActivateCommand(event) {
		return
	}

	quotes := config.Data.Quotes

	randomQuote := &quotes[rand.Intn(len(quotes))]

	ctx := context.Background()

	_, err := client.SendMessage(ctx, event.Info.Chat, &waE2E.Message{
		ExtendedTextMessage: &waE2E.ExtendedTextMessage{
			Text: randomQuote,
			ContextInfo: &waE2E.ContextInfo{
				StanzaID:      &event.Info.ID,
				Participant:   utils.Ptr(event.Info.Sender.String()),
				QuotedMessage: event.Message,
			},
		},
	})

	if err != nil {
		client.Log.Errorf("Error sending message: %v", err)
	}
}
