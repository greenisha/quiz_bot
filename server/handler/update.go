package handler

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"greenisha.ru/quiz-server/model"
)

func (h Handler) HandleUpdate(QuizResultQuestionID uint, Quiz_answerID uint, user *tgbotapi.User, MessageId int64, ChatId int64) {

	_, err := h.Rest.PostAddAnswer(QuizResultQuestionID, Quiz_answerID, user.ID, user.String())
	if err != nil {
		log.Println(err.Error())
	}
	QuizResultQuestion, err := h.Rest.GetQuizResultQuestion(QuizResultQuestionID)
	if err != nil {
		log.Println(err.Error())
	}
	question := h.ConstructQuestionText(QuizResultQuestion, false)
	log.Println(question)
	msg := tgbotapi.NewEditMessageText(int64(ChatId), int(MessageId), question)
	msg.ReplyMarkup = constructAnswers(QuizResultQuestion.Quiz_question, QuizResultQuestionID)
	log.Println("message:", QuizResultQuestion)
	h.Tgbotapi.Send(msg)

}
func (h Handler) ConstructQuestionText(QuizResultQuestion model.Quiz_result_question, isFinished bool) string {
	log.Println(QuizResultQuestion)
	var sb strings.Builder
	sb.WriteString(QuizResultQuestion.Quiz_question.Question)
	sb.WriteString("\n")
	for _, v := range QuizResultQuestion.Answers {
		sb.WriteString(" " + v.UserName) //❌
		if isFinished {
			if v.Quiz_answer.IsCorrect {
				sb.WriteString("✅")
			} else {
				sb.WriteString("❌")
			}

		}
	}

	return sb.String()
}

func (h Handler) HandleFinish(resultId uint) {
	result, err := h.Rest.GetQuizResult(int(resultId))
	if err != nil {
		log.Println(err.Error())
	}
	for _, v := range result.Quiz_result_question {
		resultQuestion, err := h.Rest.GetQuizResultQuestion(v.ID)
		if err != nil {
			log.Println(err.Error())
		}
		question := h.ConstructQuestionText(resultQuestion, true)
		log.Println(question)
		msg := tgbotapi.NewEditMessageText(result.ChatId, int(resultQuestion.MessageId), question)
		h.Tgbotapi.Send(msg)
	}
}
