package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	FirstName string    `json:"name"`
	LastName  string    `json:"name"`
	Password  string    `json:"password"`
	IsAdmin   bool      `json:"isAdmin"`
	CreatedAt time.Time `json:"createdAt"`
}

func main() {
	u := &User{
		FirstName: "John",
		LastName:  "Smith",
		Password:  "myPassword",
		IsAdmin:   true,
		CreatedAt: time.Now(),
	}

	output, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}
