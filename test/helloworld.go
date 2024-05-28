package main

import (
	"fmt"
)

const (
	english      = "English"
	spanish      = "Spanish"
	french       = "French"
	englishHello = "Hello"
	spanishHello = "Hola"
	frenchHello  = "Bonjour"
)

func Hello(name string, language string) string {
	switch name {
	case "":
		name = "World"
	}

	return greetingPrefix(language) + ", " + name + "!!!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHello
	case spanish:
		prefix = spanishHello
	default:
		prefix = englishHello
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
