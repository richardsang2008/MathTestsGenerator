package response

import (
	"time"
	"github.com/richardsang2008/MathTestsGenerator/models/compositemodels"
)

type StudentInfo struct {
	Id int `json:id`
	FName string `json:"firstName"`
	MName string `json:"midName"`
	LName string `json:"lastName"`
	StudentId string `json:"studentId""`
	Email string `json:"email"`
	EnrollmentDate time.Time `json:"enrollmentDate"`
}


type QuizItem struct {
	Id int
	LeftOperand float64
	RightOperand float64
	Answer float64
	QuizId int
	Operator compositemodels.Op

}
type Quiz struct {
	Id int `json:"id"`
	Score float64 `json:"score"`
	QuizDate time.Time `json:"quizDate"`
	Student StudentInfo `json:"student"`
	QuizItems []QuizItem `json:"quizItems"`
}
