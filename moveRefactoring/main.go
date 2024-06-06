package main

import (
	"fmt"
)

// user represents a user in the system.
type user struct {
	Name string
}

func main() {
	user := newUser("John Doe")
	fmt.Println(user.Name)
}

// newUser creates a new user instance.
func newUser(name string) user {
	return user{Name: name}
}
