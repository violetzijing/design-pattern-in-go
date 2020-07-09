package main

import "fmt"

type Subject struct {
	Observers []Observers
	State     int
}

func (s *Subject) GetState() int {
	return s.State
}

func (s *Subject) SetState(state int) {
	s.State = state
	s.NotifyAllObserver()
}

func (s *Subject) Attach(observer Observers) {
	s.Observers = append(s.Observers, observer)
}

func (s *Subject) NotifyAllObserver() {
	for _, o := range s.Observers {
		o.Update()
	}
}

type Observers interface {
	Update()
}

type BinaryObserver struct {
	Subject *Subject
}

func NewBinaryObserver(subject *Subject) *BinaryObserver {
	observer := &BinaryObserver{
		Subject: subject,
	}
	subject.Attach(observer)
	return observer
}

func (b *BinaryObserver) Update() {
	fmt.Println("binary string: ", b.Subject.GetState())
}

type OctalObserver struct {
	Subject *Subject
}

func NewOctalObserver(subject *Subject) *OctalObserver {
	observer := &OctalObserver{
		Subject: subject,
	}
	subject.Attach(observer)
	return observer
}

func (o *OctalObserver) Update() {
	fmt.Println("octal string: ", o.Subject.GetState())
}

type HexaObserver struct {
	Subject *Subject
}

func NewHexaObserver(subject *Subject) *HexaObserver {
	observer := &HexaObserver{
		Subject: subject,
	}
	subject.Attach(observer)
	return observer
}

func (h *HexaObserver) Update() {
	fmt.Println("hexa string: ", h.Subject.GetState())
}

func main() {
	subject := &Subject{}

	NewBinaryObserver(subject)
	NewOctalObserver(subject)
	NewHexaObserver(subject)

	fmt.Println("first state change: 15")
	subject.SetState(15)
	fmt.Println("second state change: 20")
	subject.SetState(20)
}
