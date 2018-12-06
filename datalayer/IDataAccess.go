package datalayer

import (
	"github.com/jinzhu/gorm"
	"github.com/richardsang2008/MathTestsGenerator/models/dbmodels"
	"github.com/richardsang2008/MathTestsGenerator/models/compositemodels"
)

type IDataAccess interface {
	NewDataAccess( db *gorm.DB) *DataAccess
	GetStudents() []dbmodels.Student
	GetStudentByStudentId(studentId string) dbmodels.Student
	GetStudent(id int) dbmodels.Student
	GetStudentByEmail(email string) dbmodels.Student
	AddStudent(firstname string, lastname string, midname string, studentId string, email string) int
	GetQuizItems() []dbmodels.QuizItem
	GetQuizItem(id int) dbmodels.QuizItem
	GetQuizItemsByQuizId(quizId int) []dbmodels.QuizItem
	AddQuizItem(leftOperand float64, rightOperand float64, operator int, answer float64, quizId int) int
	UpdateQuizItemAnswer(id int, answer float64)
	UpdateQuizScore(id int, score float64)
	GetQuizes() []dbmodels.Quiz
	GetQuiz(id int) dbmodels.Quiz
	AddQuiz(studentId string, score float64) int
	GetQuizByStudentIdAndSomeScore(studentId string, score float64) []dbmodels.Quiz
	CreateQuizItems(quizItems []compositemodels.QuizItem) []compositemodels.QuizItem
}
