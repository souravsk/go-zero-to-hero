package main

import "fmt"

const greetings = "hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return greetings + name
}

func main() {
	fmt.Println(Hello("Shubham"))
}
