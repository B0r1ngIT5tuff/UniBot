package main

import (
	"fmt"
	"io"
	"os"

	sc "github.com/B0r3ngIt5tuff/voyageBot/scraper"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {

	// Pass the token via args
	args := os.Args
	var telenews []sc.UserNews // User news
	if args[1] == " " {
		fmt.Printf("Please specify the path to the token!!\n" +
			"Usage: ./voyageBot path/to/token\n",
		)
		os.Exit(1)
	}

	t_file, ferr := os.Open(args[1])  // It takes the token path with command-line args
	TOKEN, rerr := io.ReadAll(t_file) // Reads the token

	if ferr != nil {
		fmt.Println("An error occurred while loading the token: " + ferr.Error())
		os.Exit(1)
	}

	if rerr != nil {
		fmt.Println("An error occurred while loading the token: " + rerr.Error())
		os.Exit(1)
	}

	// Create the bot with the API token
	bot, err := telego.NewBot(string(TOKEN))
	if err != nil {
		fmt.Println("Something went wrong: " + err.Error())
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

			"Ciao !! Di seguito ci sono una serie di comandi utili:\n\n"+
				" /help -> Visualizza questo messaggio\n"+
				" /news -> Trova voli economici inserendo le informazioni richieste\n"+
				" /set_news_refresh -> Imposta ogni quanto bisogna controllare per nuove notizie (TODO!)"),
		)
		if err != nil {
			bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "C'è stato un errore 😓"))
			os.Exit(1)
		}

	}, th.CommandEqual("start"))

	// Will match any message with command '/help'
	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		_, err := bot.SendMessage(tu.Message(tu.ID(message.Chat.ID),

			"Ciao !! Di seguito ci sono una serie di comandi utili:\n\n"+
				" /help -> Visualizza questo messaggio\n"+
				" /news -> Trova voli economici inserendo le informazioni richieste\n"+
				" /set_news_refresh -> Imposta ogni quanto bisogna controllare per nuove notizie (TODO!)"),
		)
		if err != nil {
			bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "C'è stato un errore 😓"))
			os.Exit(1)
		}

	}, th.CommandEqual("help"))

	// Will match  the command '/news'
	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {

		telenews = sc.GetNews() // Fetches the news

		for _, v := range telenews {
			// Send message to the user
			_, err := bot.SendMessage(tu.Message(tu.ID(message.Chat.ID),
				v.Title+"\n\n"+
					"https://www.univaq.it/"+v.Text,
			),
			)

			if err != nil {
				bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "C'è stato un errore 😓"))
				fmt.Println(err)
				os.Exit(1)
			}
		}

	}, th.CommandEqual("news"))

	bh.Start() // Starts listening

}
