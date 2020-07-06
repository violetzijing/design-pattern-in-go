package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleton map[string]string

var (
	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}

func main() {
	instance = New()
	instance["test"] = "233"
	fmt.Println(instance)

	instance = New()
	fmt.Println(instance)
}
