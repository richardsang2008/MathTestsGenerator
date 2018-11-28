package datalayer

import (
	"github.com/jinzhu/gorm"
	"github.com/richardsang2008/MathTestsGenerator/models/compositemodels"
	"github.com/richardsang2008/MathTestsGenerator/models/dbmodels"
	"math/rand"
	"time"
)

type DataAccess struct {
	db gorm.DB
}

func (r *DataAccess) GetStudents() []dbmodels.Student {
	var students []dbmodels.Student
	r.db.Find(&students)
	return students
}

func (r *DataAccess) GetStudent(id int) dbmodels.Student {
	student := dbmodels.Student{}
	r.db.Where("Id = ?", id).First(&student)
	return student
}

func (r *DataAccess) GetStudentByStudentId(studentId string) dbmodels.Student {
	student := dbmodels.Student{}
	r.db.Where("StudentId = ?", studentId).First(&student)
	return student
}

func (r *DataAccess) GetStudentByEmail(email string) dbmodels.Student {
	student := dbmodels.Student{}
	r.db.Where("Email = ?", email).First(&student)
	return student
}
func (r *DataAccess) AddStudent(firstname string, lastname string, midname string, studentId string, email string) int {
	student := dbmodels.Student{FirstName: firstname, LastName: lastname, MidName: midname, StudentId: studentId, Email: email, EnrollmentDate: time.Now()}
	r.db.Create(&student)
	return student.Id
}
func (r *DataAccess) GetQuizItems() []dbmodels.QuizItem {
	var quizItems []dbmodels.QuizItem
	r.db.Find(&quizItems)
	return quizItems
}
func (r *DataAccess) GetQuizItem(id int) dbmodels.QuizItem {
	quizItem := dbmodels.QuizItem{}
	r.db.Where("Id = ?", id).First(&quizItem)
	return quizItem
}
func (r *DataAccess) AddQuizItem(leftOperand float64, rightOperand float64, operator compositemodels.Op, answer float64, quizId int) int {
	operatorInt := 0
	if operator == compositemodels.OpADDITION {
		operatorInt = 1
	} else if operator == compositemodels.OpSUBTRACTION {
		operatorInt = 2
	} else if operator == compositemodels.OpMULTIPLICATION {
		operatorInt = 3
	} else {
		operatorInt = 4
	}
	quizItem := dbmodels.QuizItem{LeftOperand: leftOperand, RightOperand: rightOperand, Operator: operatorInt, Answer: answer, QuizId: quizId}
	r.db.Create(&quizItem)
	return quizItem.Id
}

func (r *DataAccess) UpdateQuizItemAnswer(id int, answer float64) {
	quizItem := dbmodels.QuizItem{}
	r.db.Where("Id = ?", id).First(&quizItem)
	r.db.Model(&quizItem).UpdateColumn("Answer", answer)
}
func (r *DataAccess) GetQuizes() []dbmodels.Quiz {
	var quizes []dbmodels.Quiz
	r.db.Find(&quizes)
	return quizes
}
func (r *DataAccess) GetQuiz(id int) dbmodels.Quiz {
	quiz := dbmodels.Quiz{}
	r.db.Where("Id =?", id).First(&quiz)
	return quiz
}
func (r *DataAccess) AddQuiz(studentId string, score float64) int {
	quiz := dbmodels.Quiz{StudentId: studentId, Score: score, QuizDate: time.Now()}
	r.db.Create(&quiz)
	return quiz.Id
}
func (r *DataAccess) CreateQuizItems(quizItems []compositemodels.QuizItem) {
	for _, item := range quizItems {
		r.AddQuizItem(item.LeftOperand, item.RightOperand, item.Operator, item.Answer, item.QuizId)
	}
}
func CreateQuizItem(operator compositemodels.Op, quizId int) compositemodels.QuizItem {
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
func (r *DataAccess) GenerateAQuiz(operator compositemodels.Op, studentId string) compositemodels.Quiz {
	studentdb := r.GetStudentByStudentId(studentId)
	student := compositemodels.Student{FirstName: studentdb.FirstName, MidName: studentdb.MidName, LastName: studentdb.LastName}

	retQuiz := compositemodels.Quiz{Id: 0, QuizDate: time.Now(), Score: 0, Student: student, QuizItems: []compositemodels.QuizItem{}}
	quizId := r.AddQuiz(student.StudentId, 0)
	quizItems := []compositemodels.QuizItem{}

	for i := 0; i < 10; i++ {
		quizItems = append(quizItems, CreateQuizItem(operator, quizId))
	}
	r.CreateQuizItems(quizItems)
	//map quiz
	retQuiz.Id = quizId
	retQuiz.QuizDate = time.Now()
	retQuiz.Student = student
	retQuiz.QuizItems = quizItems
	return retQuiz
}
