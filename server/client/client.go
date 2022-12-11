package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"greenisha.ru/quiz-server/model"
	"greenisha.ru/quiz-server/request"
)

type Client struct {
	RestEndpoint string
}

func getJson(url string, target interface{}) error {
	myClient := &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
func postJson(url string, target interface{}, data interface{}) error {
	myClient := &http.Client{Timeout: 10 * time.Second}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	r, err := myClient.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func (c *Client) GetQuizByName(name string) (model.Quiz, error) {
	var out model.Quiz
	err := getJson(c.RestEndpoint+"quiz?name="+name, &out)
	if err != nil {
		return model.Quiz{}, err
	}
	return out, nil
}

func (c *Client) PostQuizBegin(QuizID uint, ChatId int64) (model.Quiz_result, error) {
	request := request.QuizBegin{QuizID: QuizID, ChatId: ChatId}
	var out model.Quiz_result
	err := postJson(c.RestEndpoint+"quizBegan", &out, request)
	if err != nil {
		return model.Quiz_result{}, err
	}
	return out, nil
}
func (c *Client) GetQuizResultQuestion(id uint) (model.Quiz_result_question, error) {
	var out model.Quiz_result_question
	err := getJson(c.RestEndpoint+"quizResultQuestionByID?ID="+strconv.Itoa(int(id)), &out)
	if err != nil {
		return model.Quiz_result_question{}, err
	}
	return out, nil
}
func (c *Client) GetQuizResult(id int) (model.Quiz_result, error) {
	var out model.Quiz_result
	err := getJson(c.RestEndpoint+"quizResultByID?ID="+strconv.Itoa(id), &out)
	if err != nil {
		return model.Quiz_result{}, err
	}
	return out, nil
}
func (c *Client) GetQuiz(id int) (model.Quiz, error) {
	var out model.Quiz
	err := getJson(c.RestEndpoint+"quizByID?ID="+strconv.Itoa(id), &out)
	if err != nil {
		return model.Quiz{}, err
	}
	return out, nil
}

func (c *Client) PostAddQuestion(Quiz_resultID uint, QuestionID int64) (model.Quiz_result_question, error) {
	request := request.AddQuestion{Quiz_resultID: Quiz_resultID, QuestionID: uint(QuestionID)}
	var out model.Quiz_result_question
	err := postJson(c.RestEndpoint+"addQuestion", &out, request)
	if err != nil {
		return model.Quiz_result_question{}, err
	}
	return out, nil
}

func (c *Client) PostUpdateResultQuestion(QuizResultQuestionID uint, MessageId int64) (model.Quiz_result_question, error) {
	request := request.UpdateResponseQuestion{QuizResultQuestionID, MessageId}
	var out model.Quiz_result_question
	err := postJson(c.RestEndpoint+"updateResponseQuestion", &out, request)
	if err != nil {
		return model.Quiz_result_question{}, err
	}
	return out, nil
}

func (c *Client) PostAddAnswer(Quiz_result_questionID uint, AnswerID uint, UserID int64, UserName string) (model.Quiz_result_answer, error) {
	request := request.AddAnswer{QuizResultQuestionID: Quiz_result_questionID, AnswerID: AnswerID, UserID: UserID, UserName: UserName}
	var out model.Quiz_result_answer
	err := postJson(c.RestEndpoint+"addAnswer", &out, request)
	if err != nil {
		log.Println(err)
		return model.Quiz_result_answer{}, err
	}
	return out, nil
}
