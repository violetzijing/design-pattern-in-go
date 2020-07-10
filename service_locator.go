package main

import "fmt"

type Service interface {
	GetName() string
	Execute()
}

type Service1 struct{}

func (s *Service1) Execute() {
	fmt.Println("Executing service1")
}

func (s *Service1) GetName() string {
	return "Service1"
}

type Service2 struct{}

func (s *Service2) Execute() {
	fmt.Println("Executing service2")
}

func (s *Service2) GetName() string {
	return "Service2"
}

type InitContext struct{}

func (i *InitContext) Lookup(jndiName string) Service {
	if jndiName == "service1" {
		fmt.Println("looking up and creating a new service1 object")
		return &Service1{}
	}
	if jndiName == "service2" {
		fmt.Println("looking up and creating a new service2 object")
		return &Service2{}
	}
	return nil
}

type Cache struct {
	Services map[string]Service
}

func NewCache() *Cache {
	return &Cache{
		Services: map[string]Service{},
	}
}

func (c *Cache) GetService(serviceName string) Service {
	if v, ok := c.Services[serviceName]; ok {
		fmt.Println("return cached service: ", serviceName)
		return v
	}
	return nil
}

func (c *Cache) AddService(name string, service Service) {
	c.Services[name] = service
}

type ServiceLocator struct {
	Cache *Cache
}

func NewServiceLocator() *ServiceLocator {
	return &ServiceLocator{
		Cache: NewCache(),
	}
}

func (s *ServiceLocator) GetService(jndiName string) Service {
	service := s.Cache.GetService(jndiName)
	if service != nil {
		return service
	}
	// create a new service1
	context := &InitContext{}
	service = context.Lookup(jndiName)
	s.Cache.AddService(jndiName, service)
	return service
}

func main() {
	locator := NewServiceLocator()
	service := locator.GetService("service1")
	service.Execute()

	service = locator.GetService("service2")
	service.Execute()

	service = locator.GetService("service1")
	service.Execute()

	service = locator.GetService("service2")
	service.Execute()
}
