package autoload

import "quote-bot/config"

func init() {
	err := config.Load()
	if err != nil {
		panic(err)
	}
}
