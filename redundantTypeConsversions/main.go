package redundantTypeConsversions

import "io/ioutil"

// Hover the mouse pointer over []byte in []byte(getData()).
// Click the []byte part in []byte(getData()), press Alt + Enter and select Delete conversion.

func main() {
	_ = ioutil.WriteFile("./out.txt", []byte(getData()), 0644)
}

func getData() []byte {
	return []byte("data")
}