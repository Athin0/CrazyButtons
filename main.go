package main

import (
	"buttons/secret"
	"fmt"
	tgbotapi "github.com/skinass/telegram-bot-api/v5"
	"log"
	"net/http"
	"os"
)

var userLastCommand = make(map[int64]string)

var (
	BotToken   = secret.BotToken
	WebhookURL = secret.WebhookURL
)
var keyBoard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/button1"),
		tgbotapi.NewKeyboardButton("/button2"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/button3"),
		tgbotapi.NewKeyboardButton("/button4"),
	),
)

var keyBoardLR = []tgbotapi.InlineKeyboardButton{
	tgbotapi.NewInlineKeyboardButtonData("left", "left"),
	tgbotapi.NewInlineKeyboardButtonData("right", "right"),
}

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
		if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			if userLastCommand[update.CallbackQuery.From.ID] == "" {
				continue
			}
			data := update.CallbackQuery.Data
			from := update.CallbackQuery.From.ID
			handleCallbackQuery(bot, from, data)
			userLastCommand[update.CallbackQuery.From.ID] = ""
			continue
		}
		if update.Message == nil {
			continue
		}

		go handleUpdate(bot, update)

	}

	return nil
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("panic: %s", p)
		}
	}()

	command := update.Message.Command()
	if command == "start" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome!")
		msg.ReplyMarkup = keyBoard
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, keyBoardLR)

	switch command {
	case "button1":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You have selected button1, now select")
		msg.ReplyMarkup = keyboard
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	case "button2":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You have selected button2, now select")
		msg.ReplyMarkup = keyboard
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	case "button3":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You have selected button3, now select")
		msg.ReplyMarkup = keyboard
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	case "button4":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You have selected button4, now select")
		msg.ReplyMarkup = keyboard
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You are lucky!!!\nNo such command!\nChoose another!")
		msg.ReplyMarkup = keyBoard
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
	userLastCommand[update.Message.Chat.ID] = command
}

func handleCallbackQuery(bot *tgbotapi.BotAPI, from int64, button string) {
	msg := ""
	last := userLastCommand[from]
	switch last {
	case "button1":
		if button == "left" {
			msg = "Едут в лифте два японца, грузин, армянин и азербайджанец.\nТут один японец говорит другому: «Эти русские все на одно лицо»."
		} else {
			msg = "В школе. - Дети, давайте знакомиться. Расскажите, например, кто у кого в семье самый старший? - У меня в семье бабушка. - А у меня дедушка. - Пра-пра-пра-пра-прабабушка. - Но это же невозможно! - Во-во-во-во-возможно."
		}
	case "button2":
		if button == "left" {
			msg = "-Здравствуйте, кем работаете?\n- Я аудитор\n- Я тогда бмвтоо"
		} else {
			msg = "Как называются люди, которые боятся членистоногих и цветов?\n\nРакомакофобы"
		}
	case "button3":
		if button == "left" {
			msg = "Что сказал медведь, когда наступил на колобка? -Вот, блин!"
		} else {
			msg = "Штирлицу в голову попала пуля. \"Разрывная\" - подумал Штирлиц пораскинув мозгами."
		}
	case "button4":
		if button == "left" {
			msg = "Идёт Ленин по парку, видит – шахматисты сидят. Подошёл и слышит, как один другому говорит:\n– Давай вторую партию\nЧерез час их нафиг расстреляли, а всё потому, что партия должна быть только одна!"
		} else {
			msg = "Штирлиц стрелял вслепую. Слепая бегала зигзагами и кричала"
		}
	default:
		return
	}
	msg2 := tgbotapi.NewMessage(from, msg)
	msg2.ReplyMarkup = keyBoard
	if _, err := bot.Send(msg2); err != nil {
		log.Panic(err)
	}
}
