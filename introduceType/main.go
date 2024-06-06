package main

func main() {
	type myType struct {
		Field1 string
		Field2 int
	}
	_ = myType{
		Field1: "Hello",
		Field2: 42,
	}
}
