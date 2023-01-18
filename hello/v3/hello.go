package main

import "fmt"

const StringPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return StringPrefix + name
}

func main() {
	fmt.Println(Hello("chris"))

}
