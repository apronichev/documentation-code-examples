package main

type Base struct{}
type Derived struct {
	Base
}

func (b *Base) Method() int { return 4 }

var d Derived

func main() {
	println(d.Method())
}
