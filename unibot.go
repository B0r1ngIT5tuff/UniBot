package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {

	API_TOKEN, ferr := os.ReadFile("token.txt")

	if ferr != nil {
		fmt.Println(ferr)
		os.Exit(1)
	}

	// Create the bot with the API token
	bot, err := telego.NewBot(string(API_TOKEN))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Initialize chat updates
	updates, _ := bot.UpdatesViaLongPulling(nil)
	defer bot.StopLongPulling()

	for update := range updates {
		// Checks if there are new messages, then performs action
		if update.Message != nil {

			// Retrieve chat ID
			chatID := update.Message.Chat.ID

			//Now it will send a message to the "chatID"
			_, err := bot.SendMessage(tu.Message(tu.ID(chatID), "Fetching..."))

			if err != nil {
				bot.SendMessage(tu.Message(tu.ID(chatID), "C'è stato un errore, termino il bot."))
			}

			// END IF
		}
		// LOOP END
	}
}
