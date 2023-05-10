package main

import (
	"buttons/secret"
	"buttons/src/handlers"
	utils "buttons/src/helpers"
	"fmt"
	tgbotapi "github.com/skinass/telegram-bot-api/v5"
	"log"
	"net/http"
	"os"
)

var (
	BotToken   = secret.BotToken
	WebhookURL = secret.WebhookURL
)

func main() {
	err := startTaskBot()
	if err != nil {
		panic(err)
	}
}
func startTaskBot() error {
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Fatalf("NewBotAPI failed: %s", err)
		return err
	}

	bot.Debug = true
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	wh, err := tgbotapi.NewWebhook(WebhookURL)
	if err != nil {
		log.Fatalf("NewWebhook failed: %s", err)
		return err
	}

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatalf("SetWebhook failed: %s", err)
		return err
	}
	var userLastCommand = utils.NewLastCommand()

	updates := bot.ListenForWebhook("/")

	http.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("all is working"))
		if err != nil {
			log.Println("find state err: ", err.Error())
		}
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "80"
	}
	go func() {
		log.Fatalln("http err:", http.ListenAndServe(":"+port, nil))
	}()
	fmt.Println("start listen :" + port)

	for update := range updates {
		fmt.Println(update)
		if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			if userLastCommand.Get(update.CallbackQuery.From.ID) == "" {
				continue
			}
			data := update.CallbackQuery.Data
			from := update.CallbackQuery.From.ID
			handlers.HandleCallbackQuery(bot, from, data, userLastCommand)
			continue
		}
		if update.Message == nil {
			continue
		}
		go handlers.HandleUpdateMessage(bot, update, userLastCommand)

	}

	return nil
}

