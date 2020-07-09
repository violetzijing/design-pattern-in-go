package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Container interface {
	GetIterator() Iterator
}

type NameRepository struct {
	NameIterator *NameIterator
}

type NameIterator struct {
	Names []string
	Index int
}

func (n *NameRepository) GetIterator() Iterator {
	return n.NameIterator
}

func (n *NameIterator) HasNext() bool {
	return n.Index < len(n.Names)
}

func (n *NameIterator) Next() interface{} {
	if n.HasNext() {
		val := n.Names[n.Index]
		n.Index++
		return val
	}
	return nil
}

func main() {
	nameRepository := &NameRepository{
		NameIterator: &NameIterator{
			Names: []string{"Robert", "John", "Julie", "Lora"},
		},
	}
	for nameRepository.NameIterator.HasNext() {
		fmt.Println(nameRepository.NameIterator.Next())
	}
}
