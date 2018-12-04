package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/richardsang2008/MathTestsGenerator/models/compositemodels"
	"github.com/richardsang2008/MathTestsGenerator/models/requests"

	"github.com/jinzhu/gorm"
	"github.com/richardsang2008/MathTestsGenerator/repositories"
	"net/http"
	"strconv"
	"go.uber.org/zap"
	"github.com/gin-gonic/gin/json"
	"fmt"
)

type QuizController struct {
	Repository *repositories.Repository
	log *zap.Logger
}

func (r *QuizController) NewQuizController(l *gorm.DB, log *zap.Logger) *QuizController {
	a := repositories.Repository{}
	r.Repository = a.NewRepository(l)
	r.log = log
	return r
}
func (r *QuizController) GetAQuizById(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be a number"})
	} else {
		quiz := r.Repository.GetAQuiz(id)
		if quiz.Id == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "no record find"})
		} else {
			c.JSON(200, quiz)
		}
	}
}
func (r *QuizController) ScoreTheQuiz(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be a number"})
	} else {
		c.JSON(http.StatusOK, r.Repository.ScoreAQuiz(id))
	}
}

func (r *QuizController) AnswerAQuizItem(c *gin.Context) {
	var quizItemScore requests.QuizItemScore
	c.BindJSON(&quizItemScore)
	isValid, err := quizItemScore.IsValid()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	} else {
		if isValid == false {
			c.JSON(http.StatusBadRequest, gin.H{"message": "QuizItemId is empty"})
		} else {
			r.Repository.UpdateQuizItemAnswer(quizItemScore.QuizItemId, quizItemScore.Answer)
			c.JSON(http.StatusOK, gin.H{"message": "ok"})
		}
	}
}

func (r *QuizController) CreateAQuiz(c *gin.Context) {
	var createQuizReq requests.CreateQuiz
	c.BindJSON(&createQuizReq)
	isValid, err := createQuizReq.IsValid()
	bytes, err:= json.Marshal(createQuizReq)
	if err!= nil  {

	} else {
		r.log.Info(fmt.Sprintf("CreateAQuiz input is %s", string(bytes)))
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	} else {
		if isValid == false {
			c.JSON(http.StatusBadRequest, gin.H{"message": "StudentId is empty"})
		} else {
			//create quiz hard code to subtraction
			resp:=r.Repository.GenerateAQuiz(compositemodels.OpSUBTRACTION, createQuizReq.StudentId)
			c.JSON(http.StatusOK, resp)
			bytes, err:= json.Marshal(resp)
			if err!= nil  {

			} else {
				r.log.Info(fmt.Sprintf("CreateAQuiz out is %s", string(bytes)))
			}
		}
	}

}
