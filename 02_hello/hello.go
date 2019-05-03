package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrifix = "Hello, "
const spanishHelloPrifix = "Hola, "
const frenchHelloPrifix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrifix(language) + name
}

func greetingPrifix(language string) (prifix string) {

	switch language {
	case spanish:
		prifix = spanishHelloPrifix
	case french:
		prifix = frenchHelloPrifix
	default:
		prifix = englishHelloPrifix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
