package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return spanishPrefix + name
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
