package main

import "fmt"

type HomeView struct{}

func (h *HomeView) Show() {
	fmt.Println("displaying home page")
}

type StudentView struct{}

func (s *StudentView) Show() {
	fmt.Println("displaying student page")
}

type Dispatcher struct {
	HomeView    *HomeView
	StudentView *StudentView
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		HomeView:    &HomeView{},
		StudentView: &StudentView{},
	}
}

func (d *Dispatcher) Dispatch(request string) {
	if request == "student" {
		d.StudentView.Show()
	} else {
		d.HomeView.Show()
	}
}

type FrontController struct {
	Dispatcher *Dispatcher
}

func NewFrontController() *FrontController {
	return &FrontController{
		Dispatcher: NewDispatcher(),
	}
}

func (f *FrontController) isAuthenticUser() bool {
	fmt.Println("user is authenticated successfully")
	return true
}

func (f *FrontController) trackRequest(request string) {
	fmt.Println("Page requested: ", request)
}

func (f *FrontController) DispatherRequest(request string) {
	f.trackRequest(request)
	if f.isAuthenticUser() {
		f.Dispatcher.Dispatch(request)
	}
}

func main() {
	controller := NewFrontController()
	controller.DispatherRequest("home")
	controller.DispatherRequest("student")
}
