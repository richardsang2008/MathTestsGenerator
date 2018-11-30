package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/richardsang2008/MathTestsGenerator/models/compositemodels"
	"github.com/richardsang2008/MathTestsGenerator/models/response"
)

type IRepository interface {
	NewRepository(l *gorm.DB) *Repository
	GenerateAQuiz(operator compositemodels.Op, studentId string) compositemodels.Quiz
	ScoreAQuiz(id int) float64
	GetAQuiz(id int) response.Quiz
	UpdateQuizItemAnswer(id int, answer float64)
	AddStudent(student compositemodels.Student)  int
	GetStudentByStudentId(studentId string) response.StudentInfo
	GetStudentByEmail(email string) response.StudentInfo

}