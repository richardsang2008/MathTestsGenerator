package repositories

import "github.com/richardsang2008/MathTestsGenerator/models/compositemodels"

type IRepository interface {
	GenerateAQuiz(operator compositemodels.Op, studentId string) compositemodels.Quiz
	ScoreAQuiz(id int) float64
}