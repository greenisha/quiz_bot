package model

import (
	"html"
	"math/rand"
	"strings"
	"time"

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

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateName() string {
	rand.Seed(time.Now().UnixNano())
	for {
		name := randSeq(5)
		quiz, _ := FindQuizByName(name)
		if quiz.ID == 0 {
			return name
		}
	}
}
