package main

import (
	"fmt"
	"time"
)

type person struct {
	firstName string
	middleName string
	lastName string
	born time.Time
}

func (p person) FullName() string{
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p person) Age() int{
	return time.Now().Year() - p.born.Year()
}

func main() {
	loc, _ := time.LoadLocation("Europe/Berlin")
	newPerson := person{
		firstName: "John",
		lastName: "Smith",
		born: time.Date(2008, time.January, 1,
			23,54,00,00, loc)}
	fmt.Printf("I am %s, %s. I am %d.", newPerson.lastName, newPerson.FullName(), newPerson.Age())
}