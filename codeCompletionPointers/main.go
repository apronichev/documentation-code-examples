package main

import (
	"bytes"
)
// Place the caret between brackets of b.WriteTo. Call the SmartType action.
// Select '*myBPointerPointer' from the suggestion list.

func main() {
	b := bytes.Buffer{}
	myBPointer := &b
	myBPointerPointer := &myBPointer
	b.WriteTo()
}
