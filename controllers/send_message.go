package make_message

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
	"strconv"
)

func getToken() string {

	//tgBotToken := ""
	tgBotToken := os.Getenv("TG_BOT_TOKEN")
	log.Print("TG_BOT_TOKEN exported")

	return tgBotToken
}

func getChatId() int64 {

	chatId, err := strconv.ParseInt(os.Getenv("TG_CHAT_ID"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	//chatId := int64()
	log.Print("TG_CHAT_ID exported")

	return chatId
}

func sendMessage(text string) {

	bot, err := tgbotapi.NewBotAPI(getToken())
	if err != nil {
		log.Fatal(err)
	}

	msg := tgbotapi.NewMessage(getChatId(), text)
	msg.DisableWebPagePreview = true
	msg.ParseMode = "Markdown"
	log.Print("Message was prepared")

	_, err = bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Message was sent")
}
