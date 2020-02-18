package main

import "sync"

type S struct {
	mu sync.Mutex
}

func F() {
	r := S{}
	r = S{}
	_ = r
}

func main() {
	F()
}
