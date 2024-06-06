package main

import "fmt"

// UserService handles operations related to a user.
type UserService struct {
	// fields like database connection, etc.
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(name string) {
	fmt.Println("User created:", name)
}

// GetUser retrieves a user.
func (s *UserService) GetUser(id int) string {
	// Simulate retrieving a user
	return "User Name"
}

// UpdateUser updates an existing user.
func (s *UserService) UpdateUser(id int, newName string) {
	fmt.Println("User updated:", id, newName)
}

func main() {
	service := UserService{}
	service.CreateUser("John")
	fmt.Println("Retrieved user:", service.GetUser(1))
	service.UpdateUser(1, "Jane")
}
