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
func (p *Person) CelebrateBirthday() {
	p.Age += 1
	fmt.Printf("Happy Birthday %s! You are now %d years old.\n", p.Name, p.Age)
}

func (p *Person) IsAdult() bool {
	return p.Age >= 18
}

func main() {
	person := Person{Name: "John Doe", Age: 25, Address: []string{"123 Main St", "Anytown", "NY"}}

	person.CelebrateBirthday()

	if person.IsAdult() {
		fmt.Printf("%s is an adult.\n", person.Name)
	} else {
		fmt.Printf("%s is not an adult.\n", person.Name)
	}
}
