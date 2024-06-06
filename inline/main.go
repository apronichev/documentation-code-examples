package main

import "fmt"

func main() {
	fmt.Println(hello())
	print(hello())
}

func hello() string {
	return "Hello World!"
}
