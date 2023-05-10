package handlers

import (
	utils "buttons/src/helpers"
	tgbotapi "github.com/skinass/telegram-bot-api/v5"
	"log"
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

func HandleUpdateMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, userLastCommand *utils.LastCommand) {
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
		return
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
	userLastCommand.Set(update.Message.Chat.ID, command)
}

func HandleCallbackQuery(bot *tgbotapi.BotAPI, from int64, button string, userLastCommand *utils.LastCommand) {
	msg := ""
	last := userLastCommand.Get(from)
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
		log.Println("Strange command:", last)
		return
	}
	msg2 := tgbotapi.NewMessage(from, msg)
	msg2.ReplyMarkup = keyBoard
	if _, err := bot.Send(msg2); err != nil {
		log.Panic(err)
	}
	userLastCommand.Set(from, "")
}
