package main

import "fmt"

type Order interface {
	Execute()
}

type Stock struct {
	Name     string
	Quantity int
}

func (s *Stock) Buy() {
	fmt.Println("Stock [Name:", s.Name, ",Quantity:", s.Quantity, "] bought")
}

func (s *Stock) Sell() {
	fmt.Println("Stock [Name:", s.Name, ", Quantity:", s.Quantity, "]sold")
}

type BuyStock struct {
	ABCStock *Stock
}

func NewBuyStock(stock *Stock) *BuyStock {
	return &BuyStock{
		ABCStock: stock,
	}
}

func (b *BuyStock) Execute() {
	b.ABCStock.Buy()
}

type SellStock struct {
	ABCStock *Stock
}

func NewSellStock(stock *Stock) *SellStock {
	return &SellStock{
		ABCStock: stock,
	}
}

func (s *SellStock) Execute() {
	s.ABCStock.Sell()
}

type Broker struct {
	OrderList []Order
}

func (b *Broker) TakeOrder(order Order) {
	b.OrderList = append(b.OrderList, order)
}

func (b *Broker) PlaceOrder() {
	for _, o := range b.OrderList {
		o.Execute()
	}
	b.OrderList = []Order{}
}

func main() {
	abcStock := &Stock{
		Name:     "abc",
		Quantity: 10,
	}
	buyStock := NewSellStock(abcStock)
	sellStock := NewBuyStock(abcStock)

	broker := &Broker{
		OrderList: []Order{},
	}
	broker.TakeOrder(buyStock)
	broker.TakeOrder(sellStock)

	broker.PlaceOrder()
}
