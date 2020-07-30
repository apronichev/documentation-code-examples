package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Password  string    `json:"password"`
	IsAdmin   bool      `json:"isAdmin"`
	CreatedAt time.Time `json:"createdAt"`
}

func main() {
	u := &User{
		FirstName: "John",
		LastName:  "Smith",
		Password:  "myPassword",
		IsAdmin:   false,
		CreatedAt: time.Now(),
	}
	output := createJSON(u)
	printJSON(output)
}

//Prints JSON received from the createJSON function
func printJSON(output []byte) {
	fmt.Println(string(output))
}

//Generates JSON
func createJSON(u *User) []byte {
	output, err := json.MarshalIndent(u, "", " ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return output
}
