package main

import "fmt"

const StringPrefix = "Hello, "

func Hello(name string) string {
	return StringPrefix + name
}

func main() {
	fmt.Println(Hello("chris"))

}
