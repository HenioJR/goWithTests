package main

import "fmt"

const helloPrefix = "Hello, "

func main() {
	fmt.Printf(Hello("World"))
}

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}
