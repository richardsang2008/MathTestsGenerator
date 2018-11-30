package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/richardsang2008/MathTestsGenerator/models/requests"
	"github.com/richardsang2008/MathTestsGenerator/models/response"
	"math/rand"
	"net/http"
	"strings"
	"github.com/richardsang2008/MathTestsGenerator/repositories"
	"github.com/richardsang2008/MathTestsGenerator/models/compositemodels"
)

type StudentController struct {
	Repository *repositories.Repository
}
func (r *StudentController) NewStudentController(l *gorm.DB) *StudentController {
	a:=repositories.Repository{}
	r.Repository = a.NewRepository(l)
	return r
}
func generateRandomString(length int) string {
	characters:="0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder
	size := len(characters)
	for i:=0;i<length;i++ {
		index:= rand.Intn(size)
		a:=characters[index]
		sb.WriteString(string(a))
	}
	return sb.String()
}

func (r *StudentController) CreateStudent(c *gin.Context) {
	//get the input json
	newstudent := requests.StudentInfo{}
	c.BindJSON(&newstudent)
	isValid, err := newstudent.IsValid()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": fmt.Sprintf("new student validation has error %s ", err)})
	}
	if isValid == false {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "new student validation failed %s "})
	}
	studentId:= generateRandomString(12)
	student:= compositemodels.Student{FirstName:newstudent.FName,MidName:newstudent.MName,LastName:newstudent.LName,Email:newstudent.Email,StudentId:studentId}
	sid:=r.Repository.AddStudent(student)
	if sid >0{
		retStudent:= response.StudentInfo{FName:student.FirstName,MName:student.MidName,LName:student.LastName,
			Email:student.Email,StudentId:student.StudentId,Id:sid, EnrollmentDate:student.EnrollmentDate}
		c.JSON(http.StatusOK, retStudent)
	} else{
		c.JSON(http.StatusNotFound,gin.H{"message":"Email is used"})
	}
}

func (r *StudentController) GetStudentByStudentId(c *gin.Context) {
	studentId := c.DefaultQuery("studnetId", "NULL")
	if studentId != "NULL" && studentId != "\"\"" {
		student:=r.Repository.GetStudentByStudentId(studentId)
		if student.Email=="" {
			c.JSON(http.StatusNotFound,student)
		} else {
			c.JSON(http.StatusOK, student)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "studnetId is empty"})
	}

}

func (r *StudentController) GetStudentByEmail(c *gin.Context) {
	email := c.DefaultQuery("email", "NULL")
	if email != "NULL" && email != "\"\"" {
		student:=r.Repository.GetStudentByEmail(email)
		if student.Email=="" {
			c.JSON(http.StatusNotFound,student)
		} else {
			c.JSON(http.StatusOK, student)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "email is empty"})
	}
}
