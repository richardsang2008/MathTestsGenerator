package datalayer
import (
	"github.com/jinzhu/gorm"
	"github.com/richardsang2008/MathTestsGenerator/models/dbmodels"
	"time"
	"github.com/richardsang2008/MathTestsGenerator/models/compositemodels"
)

type DataAccess struct {
	db gorm.DB
}
func (r *DataAccess) GetStudents() []dbmodels.Student {
	var students [] dbmodels.Student
	r.db.Find(&students)
	return students
}

func (r *DataAccess) GetStudent(id int) dbmodels.Student {
	student:=dbmodels.Student{}
	r.db.Where("Id = ?",id).First(&student)
	return student
}

func (r *DataAccess) GetStudentByStudentId(studentId string) dbmodels.Student {
	student:=dbmodels.Student{}
	r.db.Where("StudentId = ?",studentId).First(&student)
	return student
}

func (r *DataAccess) GetStudentByEmail(email string) dbmodels.Student {
	student:=dbmodels.Student{}
	r.db.Where("Email = ?", email).First(&student)
	return student
}
func (r *DataAccess) AddStudent(firstname string, lastname string, midname string, studentId string, email string ) int {
	student:=dbmodels.Student{FirstName:firstname,LastName:lastname,MidName:midname,StudentId:studentId,Email:email,EnrollmentDate:time.Now()}
	r.db.Create(&student)
	return student.Id
}
func (r *DataAccess) GetQuizItems() []dbmodels.QuizItem {
	var quizItems [] dbmodels.QuizItem
	r.db.Find(&quizItems)
	return quizItems
}
func (r *DataAccess) GetQuizItem(id int) dbmodels.QuizItem {
	quizItem:=dbmodels.QuizItem{}
	r.db.Where("Id = ?", id).First(&quizItem)
	return quizItem
}
func (r *DataAccess) AddQuizItem(leftOperand float64, rightOperand float64, operator int, answer float64, quizId int) int {
	quizItem:=dbmodels.QuizItem{LeftOperand:leftOperand,RightOperand:rightOperand,Operator:operator,Answer:answer,QuizId:quizId}
	r.db.Create(&quizItem)
	return quizItem.Id
}

func (r *DataAccess) UpdateQuizItemAnswer(id int, answer float64)  {
	quizItem:=dbmodels.QuizItem{}
	r.db.Where("Id = ?", id).First(&quizItem)
	r.db.Model(&quizItem).UpdateColumn("Answer", answer)
}
func (r *DataAccess) GetQuizes() []dbmodels.Quiz {
	var quizes []dbmodels.Quiz
	r.db.Find(&quizes)
	return quizes
}
func (r *DataAccess) GetQuiz(id int) dbmodels.Quiz {
	quiz:=dbmodels.Quiz{}
	r.db.Where("Id =?",id).First(&quiz)
	return quiz
}
func (r *DataAccess) AddQuiz (studentId string, score float64) int {
	quiz:=dbmodels.Quiz{StudentId:studentId,Score:score,QuizDate:time.Now()}
	r.db.Create(&quiz)
	return quiz.Id
}
func (r *DataAccess) GenerateAQuiz(studentId string) compositemodels.Quiz {
	studentdb:= r.GetStudentByStudentId(studentId)
	student := compositemodels.Student{FirstName:studentdb.FirstName,MidName:studentdb.MidName, LastName:studentdb.LastName,}
	retQuiz := compositemodels.Quiz{Id :0, QuizDate:time.Now(),Score:0, Student:student,QuizItems:[]compositemodels.QuizItem}
}

