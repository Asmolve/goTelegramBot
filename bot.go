package main

import (
	"./config"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
	"log"
)

var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/start"),
		tgbotapi.NewKeyboardButton("/pricelist")),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/order"),
		tgbotapi.NewKeyboardButton("/help"),
	),
)
var mainOrderMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Кофе"),
		tgbotapi.NewKeyboardButton("Не кофе"),
	),
)

var coffeeOrderMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Чёрный"),
		tgbotapi.NewKeyboardButton("С молоком"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Не просто кофе"),
		tgbotapi.NewKeyboardButton("Ice"),
	),
)

var darkCoffeeOrderMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Эспрессо"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Американо"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Фильтр"),
	),
)

var milkCoffeeOrderMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Капучино"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Латте"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Flat white"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Раф"),
	),
)

var notJustCoffeeOrderMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Малиновый фильтр"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Кофейный пунш"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Манговый латте"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Раф Монте-Кристо"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Довлатте"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Кофе для котика :3"),
	),
)

var iceCoffeeOrderMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Вишнёвый бамбл"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Эспрессо-тоник"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Айс каскара"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Грушевый чай"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Персиковый чай"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Кофе для котика :3"),
	),
)

var notcoffeeOrderMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Какао"),
		tgbotapi.NewKeyboardButton("Сэндвич с рыбой"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Каскара"),
		tgbotapi.NewKeyboardButton("Сэндвич с ветчиной"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Матча-латте"),
		tgbotapi.NewKeyboardButton("Сэндвич с курицей"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Бэбичино"),
		tgbotapi.NewKeyboardButton("Зефир"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Чай"),
		tgbotapi.NewKeyboardButton("Печенье"),
	),
)

var bot *tgbotapi.BotAPI

func main() {
	fmt.Println("Sup, I'm alive again:3")
	var err error
	bot, err = tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	err = config.Connect()
	if err != nil {
		panic("Database wasn't connected :c")
	}

	var updates tgbotapi.UpdatesChannel
	updates, err = bot.GetUpdatesChan(u)

	for update := range updates {
		/*if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		switch {

		case update.Message.Text == "/start":
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome, friend "+update.Message.Chat.FirstName+":3"))

		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			//msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}*/
		/*msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ReplyMarkup = mainMenu
		bot.Send(msg)*/
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				str := "Здравствуй, " + update.Message.From.FirstName + " " + update.Message.From.LastName
				str += " 👋 \nДобро пожаловать в магазин истинных любителей кофе ☕\n"
				str += "На клавиатуре ниже представлены все доступные команды\n"
				str += ""
				str += ""
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, str))
			case "pricelist":
				bot.Send(tgbotapi.NewDocumentUpload(update.Message.Chat.ID, "./config/pricelist.jpg"))
				bot.Send(tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "./config/pricelist.jpg"))
				//bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "There will be price list"))
			case "help":
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,
					"Наш магазин продаёт кофейные зёрна, для Вашего удобства добавлены следующие кнопки \n"+
						"/start - для начала работы с ботом\n"+
						"/pricelist - вывод актуального прайс-листа\n"+
						"/order - для начала закупки\n"+
						"/help - вывод помощи по командам\n"))
			case "order":
				//bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Итак, Вы пришли за отборным кофе😋 Что Вас интересует? :)"))
				orderMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Итак, Вы пришли за отборным кофе😋 Что Вас интересует? :)")
				orderMsg.ReplyMarkup = mainOrderMenu
				bot.Send(orderMsg)
			default:
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Воспользуйтесь кнопками(командами) ниже для навигации😊"))
			}
		}
		switch update.Message.Text {
		case "Кофе":
			coffeeOrderMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Итак, Вы выбрали кофе😋 Какой Вас интересует? :)")
			coffeeOrderMsg.ReplyMarkup = coffeeOrderMenu
			bot.Send(coffeeOrderMsg)
		case "Не кофе":
			notCoffeeOrderMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Итак, Вы хотите не кофе😋 Чего изволите? :)")
			notCoffeeOrderMsg.ReplyMarkup = notcoffeeOrderMenu
			bot.Send(notCoffeeOrderMsg)

		case "Чёрный":
			darkCoffeeOrderMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Итак, выбор пал на чёрный кофе😋 Какой хотите? :)")
			darkCoffeeOrderMsg.ReplyMarkup = darkCoffeeOrderMenu
			bot.Send(darkCoffeeOrderMsg)
		case "С молоком":
			milkCoffeeOrderMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Кофе с молоком - отличный выбор😋 Что Вас интересует? :)")
			milkCoffeeOrderMsg.ReplyMarkup = milkCoffeeOrderMenu
			bot.Send(milkCoffeeOrderMsg)
		case "Не просто кофе":
			notJustCoffeeOrderMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ого, любите кофе и хотите поэксперементировать?" +
				" Похвально😋 Что желаете? :)")
			notJustCoffeeOrderMsg.ReplyMarkup = notJustCoffeeOrderMenu
			bot.Send(notJustCoffeeOrderMsg)
		case "Ice":
			iceCoffeeOrderMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Любите по-холоднее? 😋 На что пал Ваш взгляд? :)")
			iceCoffeeOrderMsg.ReplyMarkup = iceCoffeeOrderMenu
			bot.Send(iceCoffeeOrderMsg)
		}
	}
}

/*
func order(update tgbotapi.Update) {
	switch update.Message.Text {
	case "Кофе":
		upd := tgbotapi.NewUpdate(0)
		upd.Timeout = 60
		answers, _ := bot.GetUpdatesChan(upd)
		var answer tgbotapi.Update
		for answer = range answers {
			switch answer.Message.Text {
			case "Чёрный":
				switch "Чёрный" {
				case "Эспрессо":
				case "Американо":
				case "Фильтр":
				}
				switch "С молоком" {
				case "Капучино":
				case "Латте":
				case "Флет уайт":
				case "Раф":
				}
				switch "Не просто кофе" {
				case "Малиновый фильтр":
				case "Кофейный пунш":
				case "Манговый латте":
				case "Раф Монте-Кристо":
				case "Довлатте":
				case "Кофе для котика":
				}
			}
		}
	}

}
*/
