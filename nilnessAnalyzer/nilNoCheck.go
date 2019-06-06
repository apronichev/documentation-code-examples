package nilnessAnalyzer

import "os"

func main() {
	f, err := os.Open("input.txt")
	f.Close()
}
