package main

import "fmt"

type Student struct {
	RollNo string
	Name   string
}

func (s *Student) GetRollNo() string {
	return s.RollNo
}

func (s *Student) SetRollNo(str string) {
	s.RollNo = str
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) SetName(name string) {
	s.Name = name
}

type StudentView struct{}

func (v *StudentView) PrintStudentDetails(student *Student) {
	fmt.Println("Student[name: ", student.Name, ", RollNo: ", student.RollNo, "]")
}

type StudentController struct {
	Model *Student
	View  *StudentView
}

func NewStudentController(model *Student, view *StudentView) *StudentController {
	return &StudentController{
		Model: model,
		View:  view,
	}
}

func (c *StudentController) SetStudentName(name string) {
	c.Model.SetName(name)
}

func (c *StudentController) GetStudentName() string {
	return c.Model.GetName()
}

func (c *StudentController) SetStudentRollNo(roll string) {
	c.Model.SetRollNo(roll)
}

func (c *StudentController) GetStudentRollNo() string {
	return c.Model.GetRollNo()
}

func (c *StudentController) UpdateView() {
	c.View.PrintStudentDetails(c.Model)
}

func main() {
	student := &Student{RollNo: "1", Name: "Robert"}
	view := &StudentView{}
	controller := NewStudentController(student, view)
	controller.UpdateView()

	controller.SetStudentName("John")
	controller.SetStudentRollNo("2")
	controller.UpdateView()
}
