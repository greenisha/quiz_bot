package request

type QuizBegin struct {
	QuizID uint
	ChatId int64
}

type AddAnswer struct {
	QuizResultQuestionID uint
	AnswerID             uint
	UserID               int64
	UserName             string
}
type AddQuestion struct {
	Quiz_resultID uint
	QuestionID    uint
}
type UpdateResponseQuestion struct {
	QuizResultQuestionID uint
	MessageId            int64
}
type UpdateResponse struct {
	QuizResultID     uint
	QuestionPosition uint
}
type ResponseButton struct {
	QuizResultQuestionID uint
	AnswerID             uint
}
