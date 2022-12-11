package model

import "gorm.io/gorm"

type Quiz struct {
	gorm.Model
	Name       string `gorm:"size:255;not null" json:"name"`
	OtherField string
	Question   []Quiz_question
	Answer     []Quiz_answer
}

type Quiz_question struct {
	gorm.Model
	Question string `gorm:"type:text" json:"question"`
	Position int
	QuizID   uint
	Answer   []Quiz_answer
}
type Quiz_answer struct {
	gorm.Model
	Answer          string `gorm:"type:text" json:"answer"`
	Position        int
	QuizID          uint
	Quiz_questionID uint
	IsCorrect       bool `gorm:"default:false"`
}

type Quiz_result struct {
	gorm.Model
	ChatId                  int64
	QuizID                  uint
	CurrentQuestionPosition uint
	Quiz                    Quiz
	Quiz_result_question    []Quiz_result_question
}

type Quiz_result_question struct {
	gorm.Model
	Quiz_resultID   uint
	Quiz_result     Quiz_result
	Quiz_questionID uint
	Quiz_question   Quiz_question
	Answers         []Quiz_result_answer
	MessageId       uint64
}

type Quiz_result_answer struct {
	gorm.Model
	Quiz_answerID          uint
	Quiz_answer            Quiz_answer
	Quiz_result_questionID uint
	Quiz_result_question   Quiz_result_question
	UserID                 int64
	UserName               string
}
