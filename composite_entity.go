package main

import "fmt"

type DependentObject1 struct {
	Data string
}

func (d *DependentObject1) SetData(data string) {
	d.Data = data
}

func (d *DependentObject1) GetData() string {
	return d.Data
}

type DependentObject2 struct {
	Data string
}

func (d *DependentObject2) SetData(data string) {
	d.Data = data
}

func (d *DependentObject2) GetData() string {
	return d.Data
}

type CoarseGrainedObject struct {
	Do1 *DependentObject1
	Do2 *DependentObject2
}

func NewCoarseGrainedObject(data1, data2 string) *CoarseGrainedObject {
	return &CoarseGrainedObject{
		Do1: &DependentObject1{},
		Do2: &DependentObject2{},
	}
}

func (c *CoarseGrainedObject) SetData(data1, data2 string) {
	c.Do1.SetData(data1)
	c.Do2.SetData(data2)
}

func (c *CoarseGrainedObject) GetData() []string {
	return []string{c.Do1.GetData(), c.Do2.GetData()}
}

type CompositeEntity struct {
	CoarseGrainedObject *CoarseGrainedObject
}

func NewCompositeEntity(data1, data2 string) *CompositeEntity {
	return &CompositeEntity{
		CoarseGrainedObject: NewCoarseGrainedObject(data1, data2),
	}
}

func (c *CompositeEntity) SetData(data1, data2 string) {
	c.CoarseGrainedObject.SetData(data1, data2)
}

func (c *CompositeEntity) GetData() []string {
	return c.CoarseGrainedObject.GetData()
}

type Client struct {
	CompositeEntity *CompositeEntity
}

func (c *Client) PrintData() {
	for i := 0; i < len(c.CompositeEntity.GetData()); i++ {
		fmt.Println("data: ", c.CompositeEntity.GetData()[i])
	}
}

func (c *Client) SetData(data1, data2 string) {
	c.CompositeEntity.SetData(data1, data2)
}

func main() {
	client := &Client{
		CompositeEntity: NewCompositeEntity("", ""),
	}
	client.SetData("test", "data")
	client.PrintData()
	client.SetData("second test", "data2")
	client.PrintData()
}
