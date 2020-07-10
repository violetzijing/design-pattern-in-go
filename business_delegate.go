package main

import "fmt"

type BusinessService interface {
	DoProcessing()
}

type EJBService struct{}

func (e *EJBService) DoProcessing() {
	fmt.Println("Processing task by invoking EJB service")
}

type JMSService struct{}

func (j *JMSService) DoProcessing() {
	fmt.Println("Processing task by invoking JMS service")
}

type BusinessLookUp struct{}

func (b *BusinessLookUp) GetBusniessService(serviceType string) BusinessService {
	if serviceType == "EJB" {
		return &EJBService{}
	}
	return &JMSService{}
}

type BusinessDelegate struct {
	LookUpService   *BusinessLookUp
	BusinessService BusinessService
	ServiceType     string
}

func (b *BusinessDelegate) SetServiceType(serviceType string) {
	b.ServiceType = serviceType
}

func (b *BusinessDelegate) DoProcessing() {
	service := b.LookUpService.GetBusniessService(b.ServiceType)
	service.DoProcessing()
}

type Client struct {
	Service BusinessService
}

func NewClient(service BusinessService) *Client {
	return &Client{
		Service: service,
	}
}

func (c *Client) DoTask() {
	c.Service.DoProcessing()
}

func main() {
	delegate := &BusinessDelegate{LookUpService: &BusinessLookUp{}}
	delegate.SetServiceType("EJB")

	client := NewClient(delegate)
	client.DoTask()

	delegate.SetServiceType("JMS")
	client.DoTask()
}
