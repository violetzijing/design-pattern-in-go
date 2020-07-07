package main

import "fmt"

type Employee struct {
	Name         string
	Dept         string
	Salary       int
	Subordinates []*Employee
}

func NewEmployee(name, dept string, salary int) *Employee {
	return &Employee{
		Name:         name,
		Dept:         dept,
		Salary:       salary,
		Subordinates: []*Employee{},
	}
}

func (e *Employee) Add(employee *Employee) {
	e.Subordinates = append(e.Subordinates, employee)
}

func (e *Employee) Remove(employee *Employee) {
	for i, item := range e.Subordinates {
		if item.Name == employee.Name {
			e.Subordinates[i] = e.Subordinates[len(e.Subordinates)-1]
			e.Subordinates = e.Subordinates[:len(e.Subordinates)-1]
			return
		}
	}
}

func (e *Employee) GetSubordinates() []*Employee {
	return e.Subordinates
}

func (e *Employee) ToString() string {
	return fmt.Sprintf("Employee: Name: %s, Dept: %s, Salary: %d", e.Name, e.Dept, e.Salary)
}

func main() {
	ceo := NewEmployee("John", "CEO", 23333)
	headSales := NewEmployee("Robert", "HeadSales", 2333)
	headMarket := NewEmployee("Mike", "HeadMarket", 2333)
	clerk1 := NewEmployee("Laura", "marketing", 233)
	clerk2 := NewEmployee("Bob", "marketing", 233)
	sales1 := NewEmployee("Richard", "sales", 233)
	sales2 := NewEmployee("Rob", "sales", 233)

	ceo.Add(headSales)
	ceo.Add(headMarket)

	headMarket.Add(clerk1)
	headMarket.Add(clerk2)

	headSales.Add(sales1)
	headSales.Add(sales2)

	queue := []*Employee{ceo}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			employee := queue[i]
			fmt.Println(employee.ToString())
			queue = append(queue, employee.Subordinates...)
		}

		queue = queue[size:]
	}
}
