package main

import "fmt"

func Adder(a, b int) (res int) {

	res = a + b
	return
}

func main() {
	a := 2
	b := 3
	fmt.Println(Adder(a, b))

}
