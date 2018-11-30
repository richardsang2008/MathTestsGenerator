package requests

import "fmt"

type StudentInfo struct {
	FName string `json:"fName"`
	MName string `json:"mName"`
	LName string `json:"lName"`
	Email string `json:"email"`
}



func (r *StudentInfo) IsValid() (bool, error) {
	if r.FName == "" || len(r.FName) == 0 {
		return false, fmt.Errorf("First name is missing")
	}
	if r.LName == "" || len(r.LName) == 0 {
		return false, fmt.Errorf("Last name is missing")
	}
	if r.Email == "" || len(r.Email) == 0 {
		return false, fmt.Errorf("Email is missing")
	}
	return true, nil
}

type QuizItemScore struct {
	QuizItemId int     `json:"quizItemId"`
	Answer     float64 `json:"answer""`
}
func (r*QuizItemScore) IsValid() (bool,error) {

	if r.QuizItemId <1 {
		return false, fmt.Errorf("QuizItemId can not be empty")
	}
	return true, nil
}
type CreateQuiz struct{
	StudentId string `json:StudentId`
}
func(r *CreateQuiz) IsValid() (bool,error) {
	if len(r.StudentId) ==0 {
		return false, fmt.Errorf("StudentId can not be empty")
	}
	return true, nil
}