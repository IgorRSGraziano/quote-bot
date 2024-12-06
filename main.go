package main

import (
	"quote-bot/whatsapp"

	_ "quote-bot/config/autoload"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	whatsapp.Start()
}
