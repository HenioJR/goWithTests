package main

import "fmt"

const helloPrefixEng = "Hello, "
const helloPrefixSpa = "Hola, "
const spanish = "Spanish"

func main() {
	fmt.Printf(Hello("World", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return helloPrefixSpa + name
	}

	return helloPrefixEng + name
}
