package handler

import (
	"encoding/json"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"greenisha.ru/quiz-server/client"
	"greenisha.ru/quiz-server/model"
	"greenisha.ru/quiz-server/request"
)

type Handler struct {
	Tgbotapi *tgbotapi.BotAPI
	Rest     client.Client
}

func (h Handler) SendStart(chatID int64, input string) {

	quiz, err := h.Rest.GetQuizByName(input)

	if err != nil {
		log.Println(err.Error())
		return
	}
	if quiz.ID == 0 {
		return
	}
	log.Println(input, quiz)

	result, err := h.Rest.PostQuizBegin(quiz.ID, chatID)
	if err != nil {
		log.Println(err.Error())
	}
	h.sendNextQuestion(result.ID, 0)
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

	msg := tgbotapi.NewMessage(quizResult.ChatId, quiz.Question[position].Question)
	msg.ReplyMarkup = constructAnswers(quiz.Question[position], resultQuestion.ID)
	outMessage, err := h.Tgbotapi.Send(msg)
	if err != nil {
		log.Println(err.Error())
	}
	h.Rest.PostUpdateResultQuestion(resultQuestion.ID, int64(outMessage.MessageID))
	if len(quiz.Question) > position+1 {
		time.AfterFunc(5*time.Second, func() { h.sendNextQuestion(resultId, position+1) })
	}
	time.AfterFunc(10*time.Second, func() { h.HandleFinish(resultId) })
}
func constructAnswers(question model.Quiz_question, resultQuestionId uint) *tgbotapi.InlineKeyboardMarkup {
	var buttons []tgbotapi.InlineKeyboardButton

	for _, value := range question.Answer {
		answerData := request.ResponseButton{QuizResultQuestionID: resultQuestionId, AnswerID: value.ID}
		jsonData, err := json.Marshal(answerData)
		if err != nil {
			log.Println(err.Error())
		}
		button := tgbotapi.NewInlineKeyboardButtonData(value.Answer, string(jsonData))
		buttons = append(buttons, button)
	}
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			buttons...,
		),
	)
	return &numericKeyboard

}
