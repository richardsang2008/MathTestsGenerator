package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/richardsang2008/MathTestsGenerator/models/compositemodels"
)

type IRepository interface {
	NewRepository(l *gorm.DB) *Repository
	GenerateAQuiz(operator compositemodels.Op, studentId string) compositemodels.Quiz
	ScoreAQuiz(id int) float64
	AddStudent(student compositemodels.Student)  int
}