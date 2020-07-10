package main

import "fmt"

type Student struct {
	Name   string
	RollNo int
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func (s *Student) GetRollNo() int {
	return s.RollNo
}

func (s *Student) SetRollNo(no int) {
	s.RollNo = no
}

type StudentDAO interface {
	GetAllStudents() []*Student
	GetStudent(int) *Student
	UpdateStudent(*Student)
}

type StudentDAOImpl struct {
	Students []*Student
}

func NewStudentDAOImpl() *StudentDAOImpl {
	return &StudentDAOImpl{
		Students: []*Student{
			&Student{Name: "Robert", RollNo: 1},
			&Student{Name: "John", RollNo: 2},
		},
	}
}

func (s *StudentDAOImpl) GetAllStudents() []*Student {
	return s.Students
}

func (s *StudentDAOImpl) GetStudent(no int) *Student {
	for _, s := range s.Students {
		if s.RollNo == no {
			return s
		}
	}

	return nil
}

func (s *StudentDAOImpl) UpdateStudent(student *Student) {
	for _, s := range s.Students {
		if s.RollNo == student.RollNo {
			s.SetName(student.Name)
			fmt.Println("Student: roll no", student.RollNo, "updated in the database")
		}
	}
}

func main() {
	var dao StudentDAO
	dao = NewStudentDAOImpl()
	for _, s := range dao.GetAllStudents() {
		fmt.Println("Student: ", s.Name, "with roll no", s.RollNo)
	}

	student := dao.GetStudent(dao.GetAllStudents()[0].RollNo)
	student.SetName("Michael")
	dao.UpdateStudent(student)

	fmt.Println(dao.GetStudent(student.RollNo))
}
