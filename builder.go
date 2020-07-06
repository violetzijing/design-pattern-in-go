package main

import "fmt"

type Item interface {
	GetName() string
	Packing() string
	GetPrice() int
}

type Packing interface {
	Pack() string
}

type Wrapper struct{}
type Bottle struct{}

func (w *Wrapper) Pack() string {
	return "Wrapper"
}

func (b *Bottle) Pack() string {
	return "Bottle"
}

type Burger struct {
	packing Packing
}

func NewBurger() *Burger {
	return &Burger{&Wrapper{}}
}

func (b *Burger) Packing() string {
	return b.packing.Pack()
}

func (b *Burger) GetName() string {
	return "burger"
}

func (b *Burger) GetPrice() int {
	return 1
}

type ColdDrink struct {
	packing Packing
}

func NewColdDrink() *ColdDrink {
	return &ColdDrink{&Bottle{}}
}

func (c *ColdDrink) Packing() string {
	return c.packing.Pack()
}

func (c *ColdDrink) GetName() string {
	return "cold drink"
}

func (c *ColdDrink) GetPrice() int {
	return 2
}

type VegBurger struct {
	Burger *Burger
}

func (v *VegBurger) GetName() string {
	return "veg burger"
}

func (v *VegBurger) GetPrice() int {
	return v.Burger.GetPrice()
}

func (v *VegBurger) Packing() string {
	return v.Burger.Packing()
}

type ChickenBurger struct {
	Burger *Burger
}

func (c *ChickenBurger) GetName() string {
	return "chicken burger"
}

func (c *ChickenBurger) GetPrice() int {
	return c.Burger.GetPrice()
}

func (c *ChickenBurger) Packing() string {
	return c.Burger.Packing()
}

type Coke struct {
	ColdDrink *ColdDrink
}

func (c *Coke) GetName() string {
	return "coke"
}

func (c *Coke) GetPrice() int {
	return 2
}

func (c *Coke) Packing() string {
	return c.ColdDrink.Packing()
}

type Pepsi struct {
	ColdDrink *ColdDrink
}

func (p *Pepsi) GetName() string {
	return "pepsi"
}

func (p *Pepsi) GetPrice() int {
	return 3
}

func (p *Pepsi) Packing() string {
	return p.ColdDrink.Packing()
}

type Meal struct {
	Items []Item
}

func (m *Meal) AddItem(item Item) {
	m.Items = append(m.Items, item)
}

func (m *Meal) GetCost() int {
	sum := 0
	for _, i := range m.Items {
		sum += i.GetPrice()
	}

	return sum
}

func (m *Meal) ShowItems() {
	for _, i := range m.Items {
		fmt.Println("item: ", i.GetName())
		fmt.Println("packing with: ", i.Packing())
		fmt.Println("price: ", i.GetPrice())
	}
}

func main() {
	meal := &Meal{Items: []Item{}}
	meal.AddItem(&VegBurger{Burger: NewBurger()})
	meal.AddItem(&ChickenBurger{Burger: NewBurger()})
	meal.AddItem(&ChickenBurger{Burger: NewBurger()})
	meal.AddItem(&Coke{ColdDrink: NewColdDrink()})
	meal.AddItem(&Pepsi{ColdDrink: NewColdDrink()})

	fmt.Println("total cost: ", meal.GetCost())
	meal.ShowItems()
}
