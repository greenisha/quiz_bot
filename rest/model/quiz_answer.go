package model

import (
	"html"
	"strings"

	"gorm.io/gorm"
)



//	func (answer *Quiz_answer) Save() (*Quiz_answer, error) {
//		err := database.Database.Create(answer).Error
//		if err != nil {
//			return &Quiz_answer{}, err
//		}
//		return answer, nil
//	}
func (answer *Quiz_answer) BeforeSave(*gorm.DB) error {
	answer.Answer = html.EscapeString(strings.TrimSpace(answer.Answer))
	return nil
}
