package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const EngStringPrefix = "Hello, "
const SpanishStringPrefix = "Hola, "
const FrenchStringPrefix = "Bonjour, "

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

// prefix: 命名返回值
func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = SpanishStringPrefix
	case french:
		prefix = FrenchStringPrefix
	default:
		prefix = EngStringPrefix

	}

	return

}

func main() {
	fmt.Println(Hello("chris", "Spanish"))

}
