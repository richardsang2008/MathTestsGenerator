package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/richardsang2008/MathTestsGenerator/models/requests"
)

type QuizController struct {
}

func (r *QuizController) GetQuizById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id": id,
	})
}
func (r *QuizController) GetQuizScoreById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id":    id,
		"score": 100,
	})
}

func (r *QuizController) PatchQuizItems(c *gin.Context) {
	var quizItemScore requests.QuizItemScore
	c.BindJSON(&quizItemScore)
	c.JSON(200, gin.H{
		"quizItemId": quizItemScore.QuizItemId,
		"answer":     quizItemScore.Answer,
	})
}
