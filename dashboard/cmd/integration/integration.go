package integration

import "fmt"

func main() {
	val := Add(1, 2)
	fmt.Println(val)
}

func Add(a, b int) int {
	return a + b
}
