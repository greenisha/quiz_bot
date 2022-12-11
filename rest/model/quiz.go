package model

import (
	"html"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"greenisha.ru/quiz/rest/database"
)


//	func (quiz *Quiz) Save() (*Quiz, error) {
//		err := database.Database.Create(quiz).Error
//		if err != nil {
//			return &Quiz{}, err
//		}
//		return quiz, nil
//	}
func (quiz *Quiz) BeforeSave(*gorm.DB) error {
	quiz.Name = html.EscapeString(strings.TrimSpace(quiz.Name))
	return nil
}

func FindQuizByName(name string) (Quiz, error) {
	var quiz Quiz
	err := database.Database.Where("Name=?", name).Preload(clause.Associations).Preload("Question.Answer").Find(&quiz).Error
	if err != nil {
		return Quiz{}, err
	}
	return quiz, nil
}

func GetQuizByID(ID int) (Quiz, error) {
	var quiz Quiz
	err := database.Database.Preload(clause.Associations).Preload("Question.Answer").First(&quiz, ID).Error
	if err != nil {
		return Quiz{}, err
	}
	return quiz, nil

}
