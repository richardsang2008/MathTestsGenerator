package dbmodels

import "time"

type Student struct {
	Id int `gorm:"primary_key"`
	FirstName string
	MidName string
	LastName string
	StudentId string
	EnrollmentDate time.Time
	Email string
}
type Quiz struct {
	Id int `gorm:"primary_key"`
	StudentId string
	Score float64
	QuizDate time.Time
}
type QuizItem struct {
	Id int `gorm:"primary_key"`
	LeftOperand float64
	RightOperand float64
	Operator int
	Answer float64
	QuizId int
}