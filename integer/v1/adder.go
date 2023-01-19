package main

import "fmt"

func Adder(a, b int) int {

	return a + b
}

func main() {
	a := 2
	b := 3
	fmt.Println(Adder(a, b))

}
