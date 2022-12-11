package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"greenisha.ru/quiz/rest/database"
	"greenisha.ru/quiz/rest/model"
	"greenisha.ru/quiz/rest/request"
)

func Quiz(context *gin.Context) {

	quiz, err := model.FindQuizByName(context.Query("name"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if quiz.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	context.JSON(http.StatusOK, quiz)

}

func QuizBegan(context *gin.Context) {
	var input request.QuizBegin

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := model.Quiz_result{
		ChatId: input.ChatId,
		QuizID: input.QuizID,
	}
	if err := database.Database.Save(&result).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// var realResult model.Quiz_result
	database.Database.Preload(clause.Associations).First(&result, result.ID)
	context.JSON(http.StatusOK, result)
}

func QuizByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("ID"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := model.GetQuizByID(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, result)
}

func QuizResultByID(context *gin.Context) {
	var result model.Quiz_result
	id, err := strconv.Atoi(context.Query("ID"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.Database.Preload(clause.Associations).First(&result, id)
	context.JSON(http.StatusOK, result)

}

func QuizResultQuestionByID(context *gin.Context) {
	var result model.Quiz_result_question
	id, err := strconv.Atoi(context.Query("ID"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.Database.Preload(clause.Associations).Preload("Quiz_question.Answer").Preload("Answers.Quiz_answer").First(&result, id)
	context.JSON(http.StatusOK, result)
}

func AddAnswer(context *gin.Context) {
	var input request.AddAnswer
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	testAnswer, err := model.FindResultAnswer(input.QuizResultQuestionID, input.UserID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if testAnswer.ID != 0 {
		context.JSON(http.StatusConflict, testAnswer)
		return
	}

	answer := model.Quiz_result_answer{
		Quiz_answerID:          uint(input.AnswerID),
		Quiz_result_questionID: uint(input.QuizResultQuestionID),
		UserID:                 input.UserID,
		UserName:               input.UserName,
	}
	if err := database.Database.Save(&answer).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.Database.Preload(clause.Associations).First(&answer, answer.ID)
	context.JSON(http.StatusOK, answer)

}

func AddQuestion(context *gin.Context) {

	var input request.AddQuestion
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	questionResult := model.Quiz_result_question{Quiz_resultID: input.Quiz_resultID, Quiz_questionID: input.QuestionID}
	if err := database.Database.Save(&questionResult).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.Database.Preload(clause.Associations).First(&questionResult, questionResult.ID)
	context.JSON(http.StatusOK, questionResult)

}

func UpdateResponse(context *gin.Context) {
	var input request.UpdateResponse
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var curQuestion model.Quiz_result
	err := database.Database.First(&curQuestion, input.QuizResultID).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	curQuestion.CurrentQuestionPosition = input.QuestionPosition
	if err := database.Database.Save(curQuestion).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, curQuestion)

}

func UpdateResponseQuestion(context *gin.Context) {
	var input request.UpdateResponseQuestion
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var curQuestion model.Quiz_result_question
	err := database.Database.First(&curQuestion, input.QuizResultQuestionID).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	curQuestion.MessageId = uint64(input.MessageId)
	if err := database.Database.Save(curQuestion).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, curQuestion)
}
