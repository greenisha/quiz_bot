package handler

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"greenisha.ru/quiz-server/model"
	"greenisha.ru/quiz-server/request"
)

func userString(u gotgbot.User) string {
	if len(u.Username) > 0 {
		return u.Username
	} else {
		return u.FirstName + " " + u.LastName
	}
}

func (h Handler) HandleUpdate() error {
	//QuizResultQuestionID uint, Quiz_answerID uint, user *tgbotapi.User, MessageId int64, ChatId int64
	var inputData request.ResponseButton
	err := json.Unmarshal([]byte(h.Context.CallbackQuery.Data), &inputData)
	if err != nil {
		log.Println(h.Context.CallbackQuery.Data)
		panic(err)
	}
	err = h.Rest.PostAddAnswer(inputData.QuizResultQuestionID, inputData.AnswerID, h.Context.EffectiveUser.Id, userString(*h.Context.EffectiveUser))
	if err != nil {
		log.Println(err.Error())
	}
	QuizResultQuestion, err := h.Rest.GetQuizResultQuestion(inputData.QuizResultQuestionID)
	if err != nil {
		log.Println(err.Error())
	}
	question := h.ConstructQuestionText(QuizResultQuestion, false)
	log.Println(question)
	_, _, err = h.Context.EffectiveMessage.EditText(h.Bot, question, &gotgbot.EditMessageTextOpts{
		ReplyMarkup: constructAnswers(QuizResultQuestion.Quiz_question, inputData.QuizResultQuestionID),
	})
	return err

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

func (h Handler) HandleFinish(resultId uint) error {
	result, err := h.Rest.GetQuizResult(int(resultId))
	if err != nil {
		return err
	}
	for _, v := range result.Quiz_result_question {
		resultQuestion, err := h.Rest.GetQuizResultQuestion(v.ID)
		if err != nil {
			return err
		}
		question := h.ConstructQuestionText(resultQuestion, true)
		log.Println(question)
		_, _, err = h.Bot.EditMessageText(question, &gotgbot.EditMessageTextOpts{
			ChatId:    result.ChatId,
			MessageId: int64(resultQuestion.MessageId),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
