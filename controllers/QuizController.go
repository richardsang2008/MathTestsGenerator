package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/richardsang2008/MathTestsGenerator/models/requests"
	"github.com/richardsang2008/MathTestsGenerator/repositories"
	"github.com/jinzhu/gorm"
	"strconv"
	"net/http"
)

type QuizController struct {
	Repository *repositories.Repository
}
func (r *QuizController) NewQuizController(l *gorm.DB) *QuizController {
	a:=repositories.Repository{}
	r.Repository = a.NewRepository(l)
	return r
}
func (r *QuizController) GetAQuizById(c *gin.Context) {
	ids := c.Param("id")
	id,err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"id must be a number"})
	} else {
		quiz:=r.Repository.GetAQuiz(id)
		if quiz.Id ==0 {
			c.JSON(http.StatusNotFound,gin.H{"message":"no record find"})
		} else {
			c.JSON(200, quiz)
		}
	}
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
