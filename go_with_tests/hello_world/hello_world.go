package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	// printing to stdout is a side effect
	// this tutorial suggests separating side effects from pure functions
	fmt.Println(Hello("Dave"))
}
