package main

import "fmt"

type Person struct {
	ID int
	Name string
}

func (p *Person) String()  {
	fmt.Printf("Person:\n\tID:\t%d\n\tName:t%s\n", p.ID, p.Name)
}

func NewPerson (name string, id int) *Person  {
	return &Person{ID:id, Name:name}
}

func main()  {
	person := NewPerson("Admin", 1)
	fmt.Printf("person: %v\n", person)
}