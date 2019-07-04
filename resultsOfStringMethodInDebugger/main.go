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
	// Place a breakpoint on the following println function, run a debugging session.
	// See how the result is displayed in the Debug tool window (the Variables pane).
	println(u.String())
}
