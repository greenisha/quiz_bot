package model

import (
	"html"
	"strings"

	"gorm.io/gorm"
)



//	func (question *Quiz_question) Save() (*Quiz_question, error) {
//		err := database.Database.Create(question).Error
//		if err != nil {
//			return &Quiz_question{}, err
//		}
//		return question, nil
//	}
func (question *Quiz_question) BeforeSave(*gorm.DB) error {
	question.Question = html.EscapeString(strings.TrimSpace(question.Question))
	return nil
}
