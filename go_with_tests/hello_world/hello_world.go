package main

import "fmt"

const englishPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishPrefix + name + "!"
}

func main() {
	// printing to stdout is a side effect
	// this tutorial suggests separating side effects from pure functions
	fmt.Println(Hello("Dave"))
}
