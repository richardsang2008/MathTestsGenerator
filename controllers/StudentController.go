package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richardsang2008/MathTestsGenerator/models/requests"
	"github.com/richardsang2008/MathTestsGenerator/models/response"
	"log"
	"net/http"
	"strconv"
)

type StudentController struct {
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

	c.JSON(200, gin.H{"id": 1})

}

func (r *StudentController) GetStudentByStudentId(c *gin.Context) {
	studnetId := c.DefaultQuery("studnetId", "NULL")
	if studnetId != "NULL" && studnetId != "\"\"" {
		id, err := strconv.Atoi(studnetId)
		if err != nil {
			log.Println("studnetId must be a number")
			c.JSON(http.StatusBadRequest, gin.H{"Error": "studnetId must be a number"})
		}
		student := response.Student{"Joy", id}
		c.JSON(200, student)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "studnetId is empty"})
	}

}

func (r *StudentController) GetStudentByEmail(c *gin.Context) {
	email := c.DefaultQuery("email", "NULL")
	if email != "NULL" && email != "\"\"" {
		studentemail := response.StudentEmail{"Joy", email}
		c.JSON(200, studentemail)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "email is empty"})
	}
}
