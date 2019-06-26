package main

import "fmt"

func main() {
	fmt.Printf("binary number: %d\n", 0b100)
	fmt.Printf("octal number: %d\n", 0o427)
	fmt.Printf("hexadecimal numbers: %g and %g\n", 0x1p2, 0x1p-2)
	fmt.Printf("number with separators: %d\n", 1_000_000)
}
