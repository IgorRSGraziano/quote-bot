package whatsapp

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func startDb() *sqlstore.Container {
	dbLog := waLog.Stdout("Database", "INFO", true)
	container, err := sqlstore.New("sqlite3", "file:session.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	return container
}

func startClient(container *sqlstore.Container) *whatsmeow.Client {
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}

	clientLog := waLog.Stdout("Client", "INFO", true)
	return whatsmeow.NewClient(deviceStore, clientLog)
}

func login(client *whatsmeow.Client) {
	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(context.Background())
		err := client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err := client.Connect()
		if err != nil {
			panic(err)
		}
	}
}

func Start() {
	container := startDb()
	client := startClient(container)
	login(client)

	client.AddEventHandler(eventHandler(client))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}
