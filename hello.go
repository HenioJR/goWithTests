package main

import "fmt"

const helloPrefixEng = "Hello, "
const helloPrefixSpa = "Hola, "
const helloPrefixFre = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func main() {
	fmt.Printf(Hello("World", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloPrefixSpa
	case french:
		prefix = helloPrefixFre
	default:
		prefix = helloPrefixEng
	}

	return
}
