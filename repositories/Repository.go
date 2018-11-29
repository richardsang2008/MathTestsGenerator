package repositories

import (
	"github.com/richardsang2008/MathTestsGenerator/datalayer"
	"github.com/richardsang2008/MathTestsGenerator/models/compositemodels"
	"github.com/richardsang2008/MathTestsGenerator/models/dbmodels"
	"math"
	"math/rand"
	"time"
)

type Repository struct {
	DataAccessObj datalayer.DataAccess
}

func createQuizItem(operator compositemodels.Op, quizId int) compositemodels.QuizItem {
	num1 := rand.Intn(10000)
	num2 := rand.Intn(10000)
	var qi compositemodels.QuizItem
	if operator == compositemodels.OpRANDOM {
		randop := rand.Intn(4)
		if randop > 0 {
			if randop == 1 {
				operator = compositemodels.OpADDITION
			} else if randop == 2 {
				operator = compositemodels.OpSUBTRACTION
			} else if randop == 3 {
				operator = compositemodels.OpMULTIPLICATION
			} else {
				operator = compositemodels.OpDIVISION
			}
		}
	}
	if num1 < num2 {
		if num1 == 0 {
			num1 = num1 + 1
		}
		qi = compositemodels.QuizItem{Answer: 0, LeftOperand: float64(num2), RightOperand: float64(num1), QuizId: quizId, Operator: operator}
	} else {
		if num2 == 0 {
			num2 = num2 + 1
		}
		qi = compositemodels.QuizItem{Answer: 0, LeftOperand: float64(num1), RightOperand: float64(num2), QuizId: quizId, Operator: operator}
	}
	return qi
}
func (r *Repository) GenerateAQuiz(operator compositemodels.Op, studentId string) compositemodels.Quiz {
	studentdb := r.DataAccessObj.GetStudentByStudentId(studentId)
	student := compositemodels.Student{FirstName: studentdb.FirstName, MidName: studentdb.MidName, LastName: studentdb.LastName}

	retQuiz := compositemodels.Quiz{Id: 0, QuizDate: time.Now(), Score: 0, Student: student, QuizItems: []compositemodels.QuizItem{}}
	quizId := r.DataAccessObj.AddQuiz(student.StudentId, 0)
	quizItems := []compositemodels.QuizItem{}

	for i := 0; i < 10; i++ {
		quizItems = append(quizItems, createQuizItem(operator, quizId))
	}
	r.DataAccessObj.CreateQuizItems(quizItems)
	//map quiz
	retQuiz.Id = quizId
	retQuiz.QuizDate = time.Now()
	retQuiz.Student = student
	retQuiz.QuizItems = quizItems
	return retQuiz
}
func turnIntOpToOperator(operator int) compositemodels.Op {
	if operator == 1 {
		return compositemodels.OpADDITION
	} else if operator == 2 {
		return compositemodels.OpSUBTRACTION
	} else if operator == 3 {
		return compositemodels.OpMULTIPLICATION
	} else {
		return compositemodels.OpDIVISION
	}
}
func roundTo2(x float64) float64 {
	return math.Round(x*100) / 100
}
func (r *Repository) ScoreAQuiz(id int) float64 {
	quiz := r.DataAccessObj.GetQuiz(id)
	emptyQuiz := dbmodels.Quiz{}
	if quiz != emptyQuiz {
		quizItems := r.DataAccessObj.GetQuizItemsByQuizId(quiz.Id)
		if len(quizItems) > 0 {
			size := len(quizItems)
			correctCount := 0
			for _, item := range quizItems {
				turnedOp := turnIntOpToOperator(item.Operator)
				if turnedOp == compositemodels.OpADDITION {
					if roundTo2(item.Answer) == roundTo2(roundTo2(item.LeftOperand)+roundTo2(item.RightOperand)) {
						correctCount++
					}
				} else if turnedOp == compositemodels.OpSUBTRACTION {
					if roundTo2(item.Answer) == roundTo2(roundTo2(item.LeftOperand)-roundTo2(item.RightOperand)) {
						correctCount++
					}
				} else if turnedOp == compositemodels.OpMULTIPLICATION {
					if roundTo2(item.Answer) == roundTo2(roundTo2(item.LeftOperand)*roundTo2(item.RightOperand)) {
						correctCount++
					}
				} else {
					if roundTo2(item.Answer) == roundTo2(roundTo2(item.LeftOperand)/roundTo2(item.RightOperand)) {
						correctCount++
					}
				}
			}

			quiz.Score = roundTo2((float64(correctCount) / (float64(size))))
			//update the quiz score
			r.DataAccessObj.UpdateQuizItemAnswer(quiz.Id, quiz.Score)
		}

	}
	return quiz.Score
}
func (r *Repository) AddStudent(student compositemodels.Student)  int {
	id:=r.DataAccessObj.AddStudent(student.FirstName,student.LastName,student.MidName,student.StudentId,student.Email)
	return id
}