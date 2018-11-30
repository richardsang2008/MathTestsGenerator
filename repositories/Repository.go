package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/richardsang2008/MathTestsGenerator/datalayer"
	"github.com/richardsang2008/MathTestsGenerator/models/compositemodels"
	"github.com/richardsang2008/MathTestsGenerator/models/dbmodels"
	"github.com/richardsang2008/MathTestsGenerator/models/response"
	"math"
	"math/rand"
	"time"
)

type Repository struct {
	DataAccessObj *datalayer.DataAccess
}
func (r *Repository) NewRepository(l *gorm.DB) *Repository {
	m := Repository{}
	d:=datalayer.DataAccess{}
	m.DataAccessObj = d.NewDataAccess( l)
	return &m
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
func (r *Repository) GenerateAQuiz(operator compositemodels.Op, studentId string) response.Quiz {
	//verify student exist
	studentdb := r.DataAccessObj.GetStudentByStudentId(studentId)
	if len(studentdb.StudentId) >0 {
		//look for any quiz that is below the score 60
		quiz:=r.DataAccessObj.GetQuizByStudentIdAndSomeScore(studentId,60)
		if len(quiz) ==0 {
			//create a quiz
			student := compositemodels.Student{FirstName: studentdb.FirstName, MidName: studentdb.MidName, Email:studentdb.Email,
				LastName: studentdb.LastName,Id:studentdb.Id,EnrollmentDate:studentdb.EnrollmentDate,StudentId:studentdb.StudentId}
			retQuiz := response.Quiz{Id: 0, QuizDate: time.Now(), Score: 0, Student: response.StudentInfo{
				Id:student.Id,LName:student.LastName,MName:student.MidName,FName:student.FirstName,Email:student.Email,EnrollmentDate:student.EnrollmentDate,
				StudentId:student.StudentId}, QuizItems: []response.QuizItem{}}
			quizId := r.DataAccessObj.AddQuiz(student.StudentId, 0)
			quizItems := []compositemodels.QuizItem{}
			for i := 0; i < 10; i++ {
				quizItems = append(quizItems, createQuizItem(operator, quizId))
			}
			items:=r.DataAccessObj.CreateQuizItems(quizItems)
			//map quiz
			retQuiz.Id = quizId
			retQuiz.QuizDate = time.Now()
			for _, item:=range items {
				retQuiz.QuizItems = append(retQuiz.QuizItems,response.QuizItem{Id:item.Id,LeftOperand:item.LeftOperand,RightOperand:item.RightOperand,
				Answer:item.Answer,Operator:item.Operator,QuizId:item.QuizId})
			}
			return retQuiz
		} else {
			//get first unfinish quiz
			return r.GetAQuiz(quiz[0].Id)
		}
	} else {
		//no student find
		return response.Quiz{}
	}



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
func (r *Repository) UpdateQuizItemAnswer(id int, answer float64) {
	r.DataAccessObj.UpdateQuizItemAnswer(id,answer)

}
func (r *Repository) GetAQuiz(id int) response.Quiz{
	quiz:=r.DataAccessObj.GetQuiz(id)
	if quiz.StudentId!="" {
		student:=r.GetStudentByStudentId(quiz.StudentId)
		retQuiz:=response.Quiz{Id:quiz.Id,QuizDate:quiz.QuizDate,Score:quiz.Score,Student:student,QuizItems:[]response.QuizItem{}}
		quizItems:=r.DataAccessObj.GetQuizItemsByQuizId(id)
		if len(quizItems) >0 {
			for _,item:=range quizItems {
				opp :=  compositemodels.OpDIVISION
				if item.Operator ==1 {
					opp = compositemodels.OpADDITION
				} else if item.Operator ==2 {
					opp = compositemodels.OpSUBTRACTION
				} else if item.Operator ==3 {
					opp = compositemodels.OpMULTIPLICATION
				} else {
					opp = compositemodels.OpDIVISION
				}
				retQuiz.QuizItems=append(retQuiz.QuizItems,response.QuizItem{Id:item.Id,LeftOperand:item.LeftOperand,RightOperand:item.RightOperand,Answer:item.Answer,
				QuizId:item.QuizId,Operator:opp})
			}

		}
		return retQuiz
	}
	return response.Quiz{}
}

func (r *Repository) AddStudent(student compositemodels.Student)  int {
	//check if email is there
	stu:=r.DataAccessObj.GetStudentByEmail(student.Email)
	if stu.Email !="" {
		return -1
	} else {
		id:=r.DataAccessObj.AddStudent(student.FirstName,student.LastName,student.MidName,student.StudentId,student.Email)
		return id
	}

}
func (r *Repository) GetStudentByStudentId(studentId string) response.StudentInfo {
	student:=r.DataAccessObj.GetStudentByStudentId(studentId)
	if student.StudentId !="" {
		retstudent:=response.StudentInfo{StudentId:student.StudentId,FName:student.FirstName,MName:student.MidName,
		LName:student.LastName,Email:student.Email,Id:student.Id, EnrollmentDate:student.EnrollmentDate}
		return retstudent
	}
	return response.StudentInfo{}
}
func (r *Repository) GetStudentByEmail(email string) response.StudentInfo {
	student:=r.DataAccessObj.GetStudentByEmail(email)
	if student.StudentId !="" {
		retstudent:=response.StudentInfo{StudentId:student.StudentId,FName:student.FirstName,MName:student.MidName,
			LName:student.LastName,Email:student.Email,Id:student.Id, EnrollmentDate:student.EnrollmentDate}
		return retstudent
	}
	return response.StudentInfo{}
}