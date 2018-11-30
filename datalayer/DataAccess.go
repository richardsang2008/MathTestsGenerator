package datalayer

import (
	"github.com/jinzhu/gorm"
	"github.com/richardsang2008/MathTestsGenerator/models/compositemodels"
	"github.com/richardsang2008/MathTestsGenerator/models/dbmodels"

	"time"
)

type DataAccess struct {
	_db *gorm.DB
}
func (r *DataAccess) NewDataAccess( db *gorm.DB) *DataAccess {
		r._db=db
		return r
}

func (r *DataAccess) GetStudents() []dbmodels.Student {
	var students []dbmodels.Student
	r._db.Find(&students)
	return students
}

func (r *DataAccess) GetStudent(id int) dbmodels.Student {
	student := dbmodels.Student{}
	r._db.Where("Id = ?", id).First(&student)
	return student
}

func (r *DataAccess) GetStudentByStudentId(studentId string) dbmodels.Student {
	student := dbmodels.Student{}
	r._db.Where("StudentId = ?", studentId).First(&student)
	return student
}

func (r *DataAccess) GetStudentByEmail(email string) dbmodels.Student {
	student := dbmodels.Student{}
	r._db.Where("Email = ?", email).First(&student)
	return student
}
func (r *DataAccess) AddStudent(firstname string, lastname string, midname string, studentId string, email string) int {
	student := dbmodels.Student{FirstName: firstname, LastName: lastname, MidName: midname, StudentId: studentId, Email: email, EnrollmentDate: time.Now()}
	r._db.Create(&student)
	return student.Id
}
func (r *DataAccess) GetQuizItems() []dbmodels.QuizItem {
	var quizItems []dbmodels.QuizItem
	r._db.Find(&quizItems)
	return quizItems
}
func (r *DataAccess) GetQuizItem(id int) dbmodels.QuizItem {
	quizItem := dbmodels.QuizItem{}
	r._db.Where("Id = ?", id).First(&quizItem)
	return quizItem
}
func (r *DataAccess) GetQuizItemsByQuizId(quizId int) []dbmodels.QuizItem {
	quizItems := []dbmodels.QuizItem{}
	r._db.Where("QuizId = ?", quizId).Find(&quizItems)
	return quizItems
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
	r._db.Create(&quizItem)
	return quizItem.Id
}

func (r *DataAccess) UpdateQuizItemAnswer(id int, answer float64) {
	quizItem := dbmodels.QuizItem{}
	r._db.Where("Id = ?", id).First(&quizItem)
	r._db.Model(&quizItem).UpdateColumn("Answer", answer)
}
func (r *DataAccess) GetQuizes() []dbmodels.Quiz {
	var quizes []dbmodels.Quiz
	r._db.Find(&quizes)
	return quizes
}
func (r *DataAccess) GetQuiz(id int) dbmodels.Quiz {
	quiz := dbmodels.Quiz{}
	r._db.Where("Id =?", id).First(&quiz)
	return quiz
}
func (r *DataAccess) AddQuiz(studentId string, score float64) int {
	quiz := dbmodels.Quiz{StudentId: studentId, Score: score, QuizDate: time.Now()}
	r._db.Create(&quiz)
	return quiz.Id
}
func (r *DataAccess) CreateQuizItems(quizItems []compositemodels.QuizItem) {
	for _, item := range quizItems {
		r.AddQuizItem(item.LeftOperand, item.RightOperand, item.Operator, item.Answer, item.QuizId)
	}
}
