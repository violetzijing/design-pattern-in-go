package main

import "fmt"

type StudentVO struct {
	Name   string
	RollNo int
}

func NewStudentVO(name string, rollNo int) *StudentVO {
	return &StudentVO{
		Name:   name,
		RollNo: rollNo,
	}
}

func (s *StudentVO) GetName() string {
	return s.Name
}

func (s *StudentVO) SetName(name string) {
	s.Name = name
}

func (s *StudentVO) GetRollNo() int {
	return s.RollNo
}

func (s *StudentVO) SetRollNo(no int) {
	s.RollNo = no
}

type StudentBO struct {
	Students []*StudentVO
}

func NewStudentBO() *StudentBO {
	return &StudentBO{
		Students: []*StudentVO{
			&StudentVO{"Robert", 1},
			&StudentVO{"John", 2},
		},
	}
}

func (s *StudentBO) Delete(student *StudentVO) {
	for i := 0; i < len(s.Students); i++ {
		if s.Students[i].RollNo == student.RollNo {
			s.Students[i] = s.Students[len(s.Students)-1]
			s.Students = s.Students[:len(s.Students)-1]
			return
		}
	}
}

func (s *StudentBO) GetAllStudents() []*StudentVO {
	return s.Students
}

func (s *StudentBO) GetStudent(no int) *StudentVO {
	for _, item := range s.Students {
		if item.RollNo == no {
			return item
		}
	}
	return nil
}

func (s *StudentBO) UpdateStudent(student *StudentVO) {
	prev := s.GetStudent(student.RollNo)
	prev.SetName(student.Name)
	fmt.Println("Update student No: ", prev.RollNo, " name to ", prev.Name)
}

func main() {
	studentBO := NewStudentBO()
	for _, item := range studentBO.GetAllStudents() {
		fmt.Println("=======")
		fmt.Println("student NO: ", item.RollNo, ", name: ", item.Name)
	}
	fmt.Println("=======")

	student := studentBO.GetAllStudents()[0]
	student.SetName("Michael")
	studentBO.UpdateStudent(student)

	student = studentBO.GetStudent(student.RollNo)
	fmt.Println("student NO: ", student.RollNo, ", name: ", student.Name)
}
