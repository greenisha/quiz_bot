package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/joho/godotenv"
	"greenisha.ru/quiz-server/handler"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//bot, err := tgbotapi.NewBotAPI(os.Getenv("API_KEY"))
	if err != nil {
		log.Print(os.Getenv("API_KEY"))
		log.Panic(err)
	}
	token := os.Getenv("API_KEY")
	b, err := gotgbot.NewBot(token, &gotgbot.BotOpts{
		Client: http.Client{},
		DefaultRequestOpts: &gotgbot.RequestOpts{
			Timeout: gotgbot.DefaultTimeout,
			APIURL:  gotgbot.DefaultAPIURL,
		},
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}
	// Create updater and dispatcher.
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		ErrorLog: nil,
		DispatcherOpts: ext.DispatcherOpts{
			// If an error is returned by a handler, log it and continue going.
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				fmt.Println("an error occurred while handling update:", err.Error())
				return ext.DispatcherActionNoop
			},
			MaxRoutines: ext.DefaultMaxRoutines,
		},
	})
	dispatcher := updater.Dispatcher

	// /start command to introduce the bot
	dispatcher.AddHandler(handlers.NewCommand("start", start))
	dispatcher.AddHandler(handlers.NewCommand("quiz", quiz))
	dispatcher.AddHandler(handlers.NewCallback(nil, startCB))

	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	fmt.Printf("%s has been started...\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
	//updates := bot.GetUpdatesChan(u)

	// for update := range updates {
	// 	// log.Println(update)
	// 	if update.Message != nil { // If we got a message
	// 		re := regexp.MustCompile(`/quiz(@/w+)? (.+)`)

	// 		if re.FindStringSubmatch(update.Message.Text)!=nil {
	// 			handler.SendStart(update.Message.Chat.ID, re.FindStringSubmatch(update.Message.Text)[2])
	// 		}
	// 		re = regexp.MustCompile(`/start(@/w+)?`)
	// 		if re.FindStringIndex(update.Message.Text)!= nil{
	// 			handler.SendHelp(update.Message.Chat.ID)
	// 		}
	// 		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// 		// msg.ReplyToMessageID = update.Message.MessageID

	// 		// bot.Send(msg)
	// 	}
	// 	if update.CallbackQuery != nil {
	// 		var inputData request.ResponseButton
	// 		err := json.Unmarshal([]byte(update.CallbackQuery.Data), &inputData)
	// 		if err != nil {
	// 			log.Println(update.CallbackQuery.Data)
	// 			panic(err)
	// 		}
	// 		handler.HandleUpdate(inputData.QuizResultQuestionID, inputData.AnswerID, update.CallbackQuery.From, int64(update.CallbackQuery.Message.MessageID), update.CallbackQuery.Message.Chat.ID)

	// 		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)

	// 		if _, err := bot.Request(callback); err != nil {
	// 			panic(err)
	// 		}

	// 	}
	// }
}

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	return handler.NewHandler(b, ctx).SendHelp()
}

func quiz(b *gotgbot.Bot, ctx *ext.Context) error {
	re := regexp.MustCompile(`/quiz(@/w+)? (.+)`)

	if re.FindStringSubmatch(ctx.EffectiveMessage.Text) != nil {
		return handler.NewHandler(b, ctx).SendStart(re.FindStringSubmatch(ctx.EffectiveMessage.Text)[2])
	}
	return nil
}
func startCB(b *gotgbot.Bot, ctx *ext.Context) error {
	return handler.NewHandler(b, ctx).HandleUpdate()
}
