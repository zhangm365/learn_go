package main

import "fmt"

func Hello(name string) string {
	return "Hello World, " + name
}

func main() {
	fmt.Println(Hello("chris"))

}
