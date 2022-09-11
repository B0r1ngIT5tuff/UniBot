package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {

	API_TOKEN, ferr := os.ReadFile("token.txt")
	//vsc := sc.UserRequest{}

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

	// Initialize chat handler
	bh, _ := th.NewBotHandler(bot, updates)

	// Stops handling updates
	defer bh.Stop()
	// Stops getting updates
	defer bot.StopLongPulling()

	// Will match any message with command '/start'
	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		_, err := bot.SendMessage(tu.Message(tu.ID(message.Chat.ID),

			"Ciao !! Di seguito ci sono una serie di comandi utili:\n\n /help -> Visualizza questo messaggio\n /voli -> Trova voli economici inserendo le informazioni richieste\n /b_and_b -> Trova B&B economici inserendo le informazioni richieste"),
		)
		if err != nil {
			bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "C'è stato un errore 😓"))
			os.Exit(1)
		}

	}, th.CommandEqual("start"))

	// Will match any message with command '/help'
	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		_, err := bot.SendMessage(tu.Message(tu.ID(message.Chat.ID),

			"Ciao !! Di seguito ci sono una serie di comandi utili:\n\n /help -> Visualizza questo messaggio\n /voli -> Trova voli economici inserendo le informazioni richieste\n /b_and_b -> Trova B&B economici inserendo le informazioni richieste"),
		)
		if err != nil {
			bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "C'è stato un errore 😓 "))
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	}, th.CommandEqual("help"))

	// Will match  the command '/voli'
	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {

		_, err := bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), ""))

		if err != nil {
			bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "C'è stato un errore 😓"))
			fmt.Println(err)
			os.Exit(1)
		}

	}, th.AnyCallbackQuery())

}
