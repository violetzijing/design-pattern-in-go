package main

import "fmt"

type AbstructGame interface {
	EndPlay()
	Init()
	StartPlay()
}

type Game struct {
	AbstructGame AbstructGame
}

func (g *Game) Play() {
	g.AbstructGame.Init()
	g.AbstructGame.StartPlay()
	g.AbstructGame.EndPlay()
}

type Cricket struct{}

func (c *Cricket) EndPlay() {
	fmt.Println("Cricket game finished.")
}

func (c *Cricket) Init() {
	fmt.Println("Cricket game inited.")
}

func (*Cricket) StartPlay() {
	fmt.Println("Cricket game start...")
}

type Football struct{}

func (*Football) EndPlay() {
	fmt.Println("Football game finished.")
}

func (*Football) Init() {
	fmt.Println("Football game inited.")
}

func (*Football) StartPlay() {
	fmt.Println("Football game start...")
}

func main() {
	cricket := &Cricket{}
	football := &Football{}
	game := &Game{cricket}
	game.Play()
	fmt.Println("=====")

	game.AbstructGame = football
	game.Play()
}
