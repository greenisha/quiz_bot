package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"greenisha.ru/quiz/rest/controller"
	"greenisha.ru/quiz/rest/database"
	"greenisha.ru/quiz/rest/model"
)

func main() {
	loadEnv()
	loadDatabase()
	Serve()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.Quiz{})
	database.Database.AutoMigrate(&model.Quiz_answer{})
	database.Database.AutoMigrate(&model.Quiz_answer{})
	database.Database.AutoMigrate(&model.Quiz_result{})
	database.Database.AutoMigrate(&model.Quiz_result_answer{})
	database.Database.AutoMigrate(&model.Quiz_result_question{})

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Serve() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/quiz", controller.Quiz)
	router.GET("/quizByID", controller.QuizByID)
	router.GET("/quizResultByID", controller.QuizResultByID)
	router.GET("/quizResultQuestionByID", controller.QuizResultQuestionByID)
	router.POST("/updateResponse", controller.UpdateResponse)
	router.POST("/quizBegan", controller.QuizBegan)
	router.POST("/addAnswer", controller.AddAnswer)
	router.POST("/addQuestion", controller.AddQuestion)
	router.POST("/updateResponseQuestion", controller.UpdateResponseQuestion)
	router.POST("/addQuiz", controller.AddQuiz)
	router.Run(":8001")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
