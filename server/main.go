package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"greenisha.ru/quiz-server/client"
	"greenisha.ru/quiz-server/handler"
	"greenisha.ru/quiz-server/request"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_KEY"))
	if err != nil {
		log.Printf(os.Getenv("API_KEY"))
		log.Panic(err)
	}

	bot.Debug = true
	localClient := client.Client{RestEndpoint: os.Getenv("RESTAPI")}
	handler := handler.Handler{Tgbotapi: bot, Rest: localClient}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// log.Println(update)
		if update.Message != nil { // If we got a message
			re := regexp.MustCompile(`/quiz (.+)`)

			// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			handler.SendStart(update.Message.Chat.ID, re.FindStringSubmatch(update.Message.Text)[1])

			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			// msg.ReplyToMessageID = update.Message.MessageID

			// bot.Send(msg)
		}
		if update.CallbackQuery != nil {
			var inputData request.ResponseButton
			err := json.Unmarshal([]byte(update.CallbackQuery.Data), &inputData)
			if err != nil {
				log.Println(update.CallbackQuery.Data)
				panic(err)
			}
			handler.HandleUpdate(inputData.QuizResultQuestionID, inputData.AnswerID, update.CallbackQuery.From, int64(update.CallbackQuery.Message.MessageID), update.CallbackQuery.Message.Chat.ID)

			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)

			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

		}
	}
}
