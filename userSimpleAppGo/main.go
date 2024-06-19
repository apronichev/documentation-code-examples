package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Address []string
}

type Location struct {
	Street string
	City   string
	State  string
}

func (p *Person) PrintPersonDetails() {
	city := p.Address[1]
	fmt.Println(p.formatDetails(city))
}

func (p *Person) formatDetails(city string) string {
	return fmt.Sprintf("Name: %s, Age: %d, City: %s", p.Name, p.Age, city)
}

const birthdayMessage = "Happy Birthday %s! You are now %d years old.\n"

func (p *Person) celebrateBirthday() {
	p.Age += 1
	newAge := p.Age
	name := p.Name
	fmt.Printf(birthdayMessage, name, newAge)
}

func (p *Person) IsAdult() bool {
	return p.Age >= 18
}

func main() {
	person := Person{Name: "John Doe", Age: 25, Address: []string{"123 Main St", "Anytown", "NY"}}

	if person.IsAdult() {
		fmt.Printf("%s is an adult.\n", person.Name)
	} else {
		fmt.Printf("%s is not an adult.\n", person.Name)
	}
}
