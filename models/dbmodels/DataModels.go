package dbmodels

import "time"

type Student struct {
	Id int `gorm:"primary_key;column:Id"`
	FirstName string `gorm:"column:FirstName"`
	MidName string `gorm:"column:MidName"`
	LastName string `gorm:"column:LastName"`
	StudentId string `gorm:"column:StudentId"`
	EnrollmentDate time.Time `gorm:"column:EnrollmentDate"`
	Email string `gorm:"column:Email"`
}
func (Student) TableName() string {
	return "Students"
}
type Quiz struct {
	Id int `gorm:"primary_key;column:Id"`
	StudentId string `gorm:"column:StudentId"`
	Score float64 `gorm:"column:Score"`
	QuizDate time.Time `gorm:"column:QuizDate"`
}
func (Quiz) TableName() string{
	return "Quizes"
}
type QuizItem struct {
	Id int `gorm:"primary_key;column:Id"`
	LeftOperand float64 `gorm:"column:LeftOperand"`
	RightOperand float64 `gorm:"column:RightOperand"`
	Operator int `gorm:"column:Operator"`
	Answer float64 `gorm:"column:Answer"`
	QuizId int `gorm:"column:QuizId"`
}
func (QuizItem) TableName() string{
	return "QuizItems"
}