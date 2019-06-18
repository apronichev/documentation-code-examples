package main

import "fmt"

type Person struct {
	name    string
	age     int
	address Address
}

type Address struct {
	street string
}

func main() {
	// Continue typing `Person` below, select `Person` from the completion list.
	// Select `Fill All Fields` from the completion list.
	p := P
	fmt.Println(p)
}
