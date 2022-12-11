package model

import (
	"gorm.io/gorm/clause"
	"greenisha.ru/quiz/rest/database"
)

func FindResultAnswer(Quiz_result_questionID uint, UserID int64) (Quiz_result_answer, error) {
	var out Quiz_result_answer
	err := database.Database.Preload(clause.Associations).Where(Quiz_result_answer{UserID: UserID, Quiz_result_questionID: Quiz_result_questionID}).Limit(1).Find(&out).Error
	if err != nil {
		return Quiz_result_answer{}, err
	}
	return out, nil
}
