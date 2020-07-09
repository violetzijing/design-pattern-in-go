package main

import "fmt"

type Memento struct {
	State string
}

func NewMemento(state string) *Memento {
	return &Memento{
		State: state,
	}
}

func (m *Memento) GetState() string {
	return m.State
}

type Originator struct {
	State   string
	Memento *Memento
}

func (o *Originator) SetState(state string) {
	o.State = state
}

func (o *Originator) GetState() string {
	return o.State
}

func (o *Originator) SaveToMemento() *Memento {
	o.Memento = NewMemento(o.State)
	return o.Memento
}

func (o *Originator) GetStateFromMemento(memento *Memento) {
	o.State = memento.State
}

type CareTaker struct {
	MementoList []*Memento
}

func (c *CareTaker) Add(state *Memento) {
	c.MementoList = append(c.MementoList, state)
}

func (c *CareTaker) Get(index int) *Memento {
	return c.MementoList[index]
}

func main() {
	originator := &Originator{}
	careTaker := &CareTaker{}
	originator.SetState("state 1")
	originator.SetState("state 2")

	careTaker.Add(originator.SaveToMemento())
	originator.SetState("state 3")
	careTaker.Add(originator.SaveToMemento())
	originator.SetState("state 4")

	fmt.Println("current state: ", originator.State)
	originator.GetStateFromMemento(careTaker.Get(0))
	fmt.Println("first saved state: ", originator.State)
	originator.GetStateFromMemento(careTaker.Get(1))
	fmt.Println("second saved state: ", originator.State)

}
