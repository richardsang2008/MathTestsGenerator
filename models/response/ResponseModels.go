package response

type Student struct {
	Student   string `json:"Student"`
	StudnetId int    `json:"StudnetId"`
}
type StudentEmail struct {
	Student string `json:"Student"`
	Email   string `json:"email"`
}
type StudentInfo struct {
	FName string `json:"fName"`
	MName string `json:"mName"`
	LName string `json:"lName"`
	Email string `json:"email"`
	StudentId string `json:studentId`
}