package main

import "fmt"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return prefix(language) + name
}

func prefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = "Bonjour, "
	case "Spanish":
		prefix = "Hola, "
	default:
		prefix = "Hello, "
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
