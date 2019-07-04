package main

// Click the myHandler error, press Alt + Enter and select the Change signature option.
// Click the "/" error, press Alt + Enter and select the Change signature option.

type HandlerFunc func(in int, out []byte)

func myHandler() {

}

func HandleFunc(path []byte, handler HandlerFunc) {
	_, _ = path, handler
}

func main() {
	HandleFunc("/", myHandler)
}
