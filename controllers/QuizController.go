package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/richardsang2008/MathTestsGenerator/models/compositemodels"

	"github.com/richardsang2008/MathTestsGenerator/models/requests"

	"github.com/jinzhu/gorm"
	"github.com/richardsang2008/MathTestsGenerator/repositories"
	"go.uber.org/zap"
	"net/http"
	"strconv"

	"fmt"
)

type QuizController struct {
	Repository *repositories.Repository
	log        *zap.Logger
}

func (r *QuizController) NewQuizController(l *gorm.DB, log *zap.Logger) *QuizController {
	a := repositories.Repository{}
	r.Repository = a.NewRepository(l)
	r.log = log
	return r
}
func (r *QuizController) GetAQuizById(c *gin.Context) {
	ids := c.Param("id")
	r.log.Info(fmt.Sprintf("GetAQuizById input %s", ids))
	id, err := strconv.Atoi(ids)
	if err != nil {
		r.log.Error(fmt.Sprintf("GetAQuizById input %s is not a number", ids))
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be a number"})
	} else {
		quiz := r.Repository.GetAQuiz(id)
		if quiz.Id == 0 {
			r.log.Info(fmt.Sprintf("GetAQuizById found no quiz by id %s", ids))
			c.JSON(http.StatusNotFound, gin.H{"message": "no record find"})
		} else {
			bytes, err := json.Marshal(quiz)
			if err != nil {
				r.log.Error(fmt.Sprintf("GetAQuizById is error out %s", err))
				c.JSON(http.StatusBadRequest, gin.H{"message": err})
			} else {
				r.log.Info(fmt.Sprintf("GetAQuizById response %s", string(bytes)))
				c.JSON(200, quiz)
			}
		}
	}
}
func (r *QuizController) ScoreTheQuiz(c *gin.Context) {
	ids := c.Param("id")
	r.log.Info(fmt.Sprintf("ScoreTheQuiz input %s", ids))
	id, err := strconv.Atoi(ids)
	if err != nil {
		r.log.Error(fmt.Sprintf("ScoreTheQuiz input %s is not a number", ids))
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be a number"})
	} else {
		resp := r.Repository.ScoreAQuiz(id)
		bytes, err := json.Marshal(resp)
		if err != nil {
			r.log.Error(fmt.Sprintf("ScoreTheQuiz is error out %s", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
		} else {
			r.log.Error(fmt.Sprintf("ScoreTheQuiz response %s", string(bytes)))
			c.JSON(http.StatusOK, r.Repository.ScoreAQuiz(id))
		}
	}
}

func (r *QuizController) AnswerAQuizItem(c *gin.Context) {
	var quizItemScore requests.QuizItemScore
	c.BindJSON(&quizItemScore)
	isValid, err := quizItemScore.IsValid()
	if err != nil {
		r.log.Error(fmt.Sprintf("AnswerAQuizItem is error out %s", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	} else {
		if isValid == false {
			r.log.Error(fmt.Sprintf("AnswerAQuizItem is not valid %s"))
			c.JSON(http.StatusBadRequest, gin.H{"message": "QuizItemId is empty"})
		} else {
			bytes, err := json.Marshal(quizItemScore)
			if err != nil {
				r.log.Error(fmt.Sprintf("AnswerAQuizItem is not valid %s"))
				c.JSON(http.StatusBadRequest, gin.H{"message": err})
			} else {
				r.log.Info(fmt.Sprintf("AnswerAQuizItem input %s", string(bytes)))
			}
			r.Repository.UpdateQuizItemAnswer(quizItemScore.QuizItemId, quizItemScore.Answer)
			r.log.Info("AnswerAQuizItem response OK")
			c.JSON(http.StatusOK, gin.H{"message": "ok"})
		}
	}
}

func (r *QuizController) CreateAQuiz(c *gin.Context) {
	var createQuizReq requests.CreateQuiz
	c.BindJSON(&createQuizReq)
	isValid, err := createQuizReq.IsValid()
	if err != nil {
		r.log.Error(fmt.Sprintf("CreateQuizReq is error out %s", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	} else {
		bytes, err := json.Marshal(createQuizReq)
		if err != nil {
			r.log.Error(fmt.Sprintf("CreateQuizReq is error out %s", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
		} else {
			if isValid == false {
				r.log.Info(fmt.Sprintf("CreateAQuiz input is not valid %s", string(bytes)))
				c.JSON(http.StatusBadRequest, gin.H{"message": "StudentId is empty"})
			} else {
				//create quiz hard code to subtraction
				r.log.Info(fmt.Sprintf("CreateAQuiz input is %s", string(bytes)))
				resp := r.Repository.GenerateAQuiz(compositemodels.OpSUBTRACTION, createQuizReq.StudentId)
				bytes, err = json.Marshal(resp)
				if err != nil {
					r.log.Error(fmt.Sprintf("CreateAQuiz response error %s", err))
					c.JSON(http.StatusInternalServerError, gin.H{"message": err})
				} else {
					r.log.Info(fmt.Sprintf("CreateAQuiz response %s", string(bytes)))
					c.JSON(http.StatusOK, resp)
				}
			}
		}
	}
}
