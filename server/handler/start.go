package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"greenisha.ru/quiz-server/client"
	"greenisha.ru/quiz-server/model"
	"greenisha.ru/quiz-server/request"
)

func NewHandler(b *gotgbot.Bot, ctx *ext.Context) Handler {
	return Handler{Bot: b, Context: ctx, Rest: client.Client{RestEndpoint: os.Getenv("RESTAPI")}}
}

type Handler struct {
	//Tgbotapi *tgbotapi.BotAPI
	Rest    client.Client
	Bot     *gotgbot.Bot
	Context *ext.Context
}

func (h Handler) SendStart(name string) error {

	//chatID int64, input string
	chatID := h.Context.EffectiveChat.Id
	input := name

	log.Println(input)
	quiz, err := h.Rest.GetQuizByName(input)

	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("failed to send start message: %w", err)
	}
	if quiz.ID == 0 {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	log.Println(input, quiz)

	result, err := h.Rest.PostQuizBegin(quiz.ID, chatID)
	if err != nil {
		log.Println(err.Error())
	}
	h.sendNextQuestion(result.ID, 0)
	return nil
}
func (h Handler) SendHelp() error {
	_, err := h.Context.EffectiveMessage.Reply(h.Bot, "Bot for making and sending quizes to chats and so on", &gotgbot.SendMessageOpts{ParseMode: "html"})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

func (h Handler) sendNextQuestion(resultId uint, position int) {
	quizResult, err := h.Rest.GetQuizResult(int(resultId))
	if err != nil {
		log.Println(err.Error())
	}
	quiz, err := h.Rest.GetQuiz(int(quizResult.QuizID))
	if err != nil {
		log.Println(err.Error())
	}

	resultQuestion, err := h.Rest.PostAddQuestion(quizResult.ID, int64(quiz.Question[position].ID))

	if err != nil {
		log.Println(err.Error())
	}

	msg, err := h.Context.EffectiveMessage.Reply(h.Bot, quiz.Question[position].Question,
		&gotgbot.SendMessageOpts{
			ParseMode:   "html",
			ReplyMarkup: constructAnswers(quiz.Question[position], resultQuestion.ID),
		})

	if err != nil {
		log.Println(err.Error())
	}
	h.Rest.PostUpdateResultQuestion(resultQuestion.ID, int64(msg.MessageId))
	if len(quiz.Question) > position+1 {
		time.AfterFunc(5*time.Second, func() { h.sendNextQuestion(resultId, position+1) })
	}
	time.AfterFunc(10*time.Second, func() { h.HandleFinish(resultId) })
}
func constructAnswers(question model.Quiz_question, resultQuestionId uint) gotgbot.InlineKeyboardMarkup {
	var buttons [][]gotgbot.InlineKeyboardButton
	rowCount := 0
	var buttonsRow []gotgbot.InlineKeyboardButton
	for _, value := range question.Answer {
		rowCount++
		answerData := request.ResponseButton{QuizResultQuestionID: resultQuestionId, AnswerID: value.ID}
		jsonData, err := json.Marshal(answerData)
		if err != nil {
			log.Println(err.Error())
		}
		button := gotgbot.InlineKeyboardButton{Text: value.Answer, CallbackData: string(jsonData)}
		buttonsRow = append(buttonsRow, button)
		if rowCount >= 2 {
			buttons = append(buttons, buttonsRow)
			buttonsRow = nil
		}
	}
	if buttonsRow != nil {
		buttons = append(buttons, buttonsRow)
	}
	var numericKeyboard = gotgbot.InlineKeyboardMarkup{InlineKeyboard: buttons}
	return numericKeyboard

}
