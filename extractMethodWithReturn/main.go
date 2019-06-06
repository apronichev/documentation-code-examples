package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter text")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "quit"{
			return
		} else if text == "boop" {
			fmt.Printf("boop!")
		} else {
			fmt.Printf("IDK")
		}
	}
}
