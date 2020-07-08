package main

import "fmt"

type Image interface {
	Display()
}

type RealImage struct {
	Filename string
}

func NewRealImage(filename string) *RealImage {
	image := &RealImage{
		Filename: filename,
	}
	image.LoadFromDisk()

	return image
}

func (r *RealImage) Display() {
	fmt.Println("display ", r.Filename)
}

func (r *RealImage) LoadFromDisk() {
	fmt.Println("loading ", r.Filename)
}

type ProxyImage struct {
	RealImage *RealImage
	Filename  string
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{
		Filename: filename,
	}
}

func (p *ProxyImage) Display() {
	if p.RealImage == nil {
		p.RealImage = NewRealImage(p.Filename)
	}
	p.RealImage.Display()
}

func main() {
	var image Image
	image = NewProxyImage("test.jpg")
	image.Display()
	fmt.Println("image doesn't need to be loaded")
	image.Display()
}
