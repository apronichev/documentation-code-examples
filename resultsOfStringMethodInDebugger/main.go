package main

import (
	"fmt"
)

type User struct {
	name    string
	email   string
	company Company
}

type Company struct {
	name string
}

func (u User) String() string {
	return fmt.Sprintf("%s <%s> (%s)", u.name, u.email, u.company.name)
}

func main() {
	u := User{
		name:    "John Smith",
		email:   "john.smith@example.com",
		company: Company{name: "ACME"},
	}
	println(u.String())
}
