package main

import "fmt"

type Filter interface {
	Execute(string)
}

type AuthenticationFilter struct{}

func (a *AuthenticationFilter) Execute(request string) {
	fmt.Println("AuthenticationFilter request: ", request)
}

type DebugFilter struct{}

func (d *DebugFilter) Execute(request string) {
	fmt.Println("DebugFilter request: ", request)
}

type Target struct{}

func (t *Target) Execute(request string) {
	fmt.Println("Executing request: ", request)
}

type FilterChain struct {
	Filters []Filter
	Target  *Target
}

func (f *FilterChain) AddFilter(filter Filter) {
	f.Filters = append(f.Filters, filter)
}

func (f *FilterChain) Execute(request string) {
	for _, f := range f.Filters {
		f.Execute(request)
	}
	f.Target.Execute(request)
}

func (f *FilterChain) SetTarget(target *Target) {
	f.Target = target
}

type FilterManager struct {
	FilterChain *FilterChain
}

func NewFilterManager(target *Target) *FilterManager {
	return &FilterManager{
		FilterChain: &FilterChain{
			Filters: []Filter{},
			Target:  target,
		},
	}
}

func (f *FilterManager) SetFilter(filter Filter) {
	f.FilterChain.AddFilter(filter)
}

func (f *FilterManager) FilterRequest(request string) {
	f.FilterChain.Execute(request)
}

type Client struct {
	FilterManager *FilterManager
}

func (c *Client) SetFilterManager(manager *FilterManager) {
	c.FilterManager = manager
}

func (c *Client) SendRequest(request string) {
	c.FilterManager.FilterRequest(request)
}

func main() {
	manager := NewFilterManager(&Target{})
	manager.SetFilter(&AuthenticationFilter{})
	manager.SetFilter(&DebugFilter{})

	client := &Client{FilterManager: manager}
	client.SendRequest("home")
}
