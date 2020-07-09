package main

import "fmt"

type AbstractCustomer interface {
	IsNil() bool
	GetName() string
}

type RealCustomer struct {
	Name string
}

func NewRealCustomer(name string) *RealCustomer {
	return &RealCustomer{
		Name: name,
	}
}

func (r *RealCustomer) GetName() string {
	return r.Name
}

func (*RealCustomer) IsNil() bool {
	return false
}

type NullCustomer struct{}

func (n *NullCustomer) GetName() string {
	return "Not available in customer db"
}

func (*NullCustomer) IsNil() bool {
	return true
}

type CustomerFactory struct {
	Names []string
}

func NewCustomerFacotry() *CustomerFactory {
	return &CustomerFactory{
		Names: []string{"Rob", "Joe", "Julie"},
	}
}

func (c *CustomerFactory) GetCustomer(name string) AbstractCustomer {
	for _, n := range c.Names {
		if n == name {
			return NewRealCustomer(name)
		}
	}
	return &NullCustomer{}
}

func main() {
	factory := NewCustomerFacotry()
	customer1 := factory.GetCustomer("Rob")
	customer2 := factory.GetCustomer("Bob")
	customer3 := factory.GetCustomer("Julie")
	customer4 := factory.GetCustomer("Laura")

	fmt.Println("customers: ")
	fmt.Println(customer1.GetName())
	fmt.Println(customer2.GetName())
	fmt.Println(customer3.GetName())
	fmt.Println(customer4.GetName())
}
