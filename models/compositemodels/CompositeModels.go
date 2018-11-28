package compositemodels

import "time"

type Student struct {
	Id int
	FirstName string
	MidName string
	LastName string
	StudentId string
	Email string
	EnrollmentDate time.Time
}
type QuizItem struct {
	Id int
	LeftOperand float64
	RightOperand float64
	Answer float64
	QuizId int

}
type Quiz struct {
	Id int
	Score float64
	QuizDate time.Time
	Student Student
	QuizItems []QuizItem
}
