package main

import "fmt"

type Person struct {
	Name         string
	Gender       string
	MarialStatus string
}

func NewPerson(name, gender, marialStatus string) *Person {
	return &Person{
		Name:         name,
		Gender:       gender,
		MarialStatus: marialStatus,
	}
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetGender() string {
	return p.Gender
}

func (p *Person) GetMarialStatus() string {
	return p.MarialStatus
}

type Criteria interface {
	MeetCriteria(people []*Person) []*Person
}

type MaleCriteria struct{}

func (m *MaleCriteria) MeetCriteria(people []*Person) []*Person {
	male := []*Person{}
	for _, p := range people {
		if p.Gender == "male" {
			male = append(male, p)
		}
	}

	return male
}

type FemaleCriteria struct{}

func (f *FemaleCriteria) MeetCriteria(people []*Person) []*Person {
	female := []*Person{}
	for _, p := range people {
		if p.Gender == "female" {
			female = append(female, p)
		}
	}

	return female
}

type SingleCriteria struct{}

func (s *SingleCriteria) MeetCriteria(people []*Person) []*Person {
	single := []*Person{}
	for _, p := range people {
		if p.MarialStatus == "single" {
			single = append(single, p)
		}
	}

	return single
}

type AndCriteria struct {
	Criteria      Criteria
	OtherCriteria Criteria
}

func (a *AndCriteria) MeetCriteria(people []*Person) []*Person {
	result1 := a.Criteria.MeetCriteria(people)
	return a.OtherCriteria.MeetCriteria(result1)
}

type OrCriteria struct {
	Criteria      Criteria
	OtherCriteria Criteria
}

func (o *OrCriteria) MeetCriteria(people []*Person) []*Person {
	result1 := o.Criteria.MeetCriteria(people)
	hash := map[string]*Person{}
	for _, p := range result1 {
		hash[p.Name] = p
	}
	result2 := o.OtherCriteria.MeetCriteria(people)
	for _, p := range result2 {
		if _, ok := hash[p.Name]; !ok {
			result2 = append(result2, p)
		}
	}

	return result2
}

func main() {
	people := []*Person{}
	people = append(people, &Person{
		Name: "1", Gender: "male", MarialStatus: "single",
	})
	people = append(people, &Person{
		Name: "2", Gender: "female", MarialStatus: "single",
	})
	people = append(people, &Person{
		Name: "3", Gender: "male", MarialStatus: "married",
	})
	people = append(people, &Person{
		Name: "4", Gender: "male", MarialStatus: "married",
	})
	people = append(people, &Person{
		Name: "5", Gender: "female", MarialStatus: "single",
	})
	people = append(people, &Person{
		Name: "6", Gender: "female", MarialStatus: "single",
	})
	people = append(people, &Person{
		Name: "7", Gender: "female", MarialStatus: "single",
	})

	male := &MaleCriteria{}
	female := &FemaleCriteria{}
	single := &SingleCriteria{}
	singleFemale := &AndCriteria{single, female}
	singleOrMale := &OrCriteria{single, male}

	fmt.Println("===== male =====")
	for _, p := range male.MeetCriteria(people) {
		fmt.Println(p)
	}
	fmt.Println("===== female =====")
	for _, p := range female.MeetCriteria(people) {
		fmt.Println(p)
	}
	fmt.Println("===== single =====")
	for _, p := range single.MeetCriteria(people) {
		fmt.Println(p)
	}
	fmt.Println("===== single female =====")
	for _, p := range singleFemale.MeetCriteria(people) {
		fmt.Println(p)
	}
	fmt.Println("===== single or male =====")
	for _, p := range singleOrMale.MeetCriteria(people) {
		fmt.Println(p)
	}
}
