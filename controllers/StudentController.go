package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"github.com/jinzhu/gorm"
	"github.com/richardsang2008/MathTestsGenerator/models/requests"
	"github.com/richardsang2008/MathTestsGenerator/models/response"
	"math/rand"
	"net/http"
	"strings"
	"github.com/richardsang2008/MathTestsGenerator/repositories"
	"github.com/richardsang2008/MathTestsGenerator/models/compositemodels"
	"go.uber.org/zap"
)

type StudentController struct {
	Repository *repositories.Repository
	log *zap.Logger
}
func (r *StudentController) NewStudentController(l *gorm.DB, log *zap.Logger) *StudentController {
	a:=repositories.Repository{}
	r.Repository = a.NewRepository(l)
	r.log = log
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
		r.log.Error(fmt.Sprintf("CreateStudent input error %s",err))
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	if isValid == false {
		r.log.Error(fmt.Sprintf("CreateStudent input is not valid"))
		c.JSON(http.StatusBadRequest, gin.H{"message": "new student validation failed %s "})
	}
	bytes,err := json.Marshal(newstudent)
	if err != nil {
		r.log.Error(fmt.Sprintf("CreateStudent input error %s",err))
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	} else {
		r.log.Info(fmt.Sprintf("CreateStudent input %s",string(bytes)))
	}
	studentId:= generateRandomString(12)
	student:= compositemodels.Student{FirstName:newstudent.FName,MidName:newstudent.MName,LastName:newstudent.LName,Email:newstudent.Email,StudentId:studentId}
	sid:=r.Repository.AddStudent(student)
	if sid >0{
		retStudent:= response.StudentInfo{FName:student.FirstName,MName:student.MidName,LName:student.LastName,
			Email:student.Email,StudentId:student.StudentId,Id:sid, EnrollmentDate:student.EnrollmentDate}
		bytes,err = json.Marshal(retStudent)
		if err != nil {
			r.log.Error(fmt.Sprintf("CreateStudent response error %s",err))
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
		} else {
			r.log.Info(fmt.Sprintf("CreateStudent response %s", string(bytes)))
			c.JSON(http.StatusOK, retStudent)
		}

	} else{
		r.log.Info("CreateStudent response email is used")
		c.JSON(http.StatusNotFound,gin.H{"message":"Email is used"})
	}
}

func (r *StudentController) GetStudentByStudentId(c *gin.Context) {
	studentId := c.DefaultQuery("studnetId", "NULL")
	r.log.Info(fmt.Sprintf("GetStudentByStudentId input %s",studentId))
	if studentId != "NULL" && studentId != "\"\"" {
		student:=r.Repository.GetStudentByStudentId(studentId)
		if student.Email=="" {
			r.log.Info(fmt.Sprintf("GetStudentByStudentId return no student by input %s", studentId))
			c.JSON(http.StatusNotFound,student)
		} else {
			r.log.Info("GetStudentByStudentId return OK")
			c.JSON(http.StatusOK, student)
		}
	} else {
		r.log.Error(fmt.Sprintf("GetStudentByStudentId input is empty %s", studentId))
		c.JSON(http.StatusBadRequest, gin.H{"Error": "studnetId is empty"})
	}

}

func (r *StudentController) GetStudentByEmail(c *gin.Context) {
	email := c.DefaultQuery("email", "NULL")
	r.log.Info(fmt.Sprintf("GetStudentByEmail input %s", email))
	if email != "NULL" && email != "\"\"" {
		student:=r.Repository.GetStudentByEmail(email)
		if student.Email=="" {
			r.log.Info(fmt.Sprintf("GetStudentByEmail return no record by email %s",email))
			c.JSON(http.StatusNotFound,student)
		} else {
			bytes, err:= json.Marshal(student)
			if err != nil {
				r.log.Error(fmt.Sprintf("GetStudentByEmail return error %s",err))
				c.JSON(http.StatusInternalServerError, gin.H{"message":err})
			} else {
				r.log.Info(fmt.Sprintf("GetStudentByEmail return %s", string(bytes)))
				c.JSON(http.StatusOK, student)
			}
		}
	} else {
		r.log.Error("GetStudentByEmail input is empty")
		c.JSON(http.StatusBadRequest, gin.H{"Error": "email is empty"})
	}
}
