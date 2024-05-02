package main

import "fmt"

const (
	romanian = "Romanian"
	french = "French"

	englishHelloPrefix = "Hello, "
	romanianHelloPrefix = "Salut, "
	frenchHelloPrefix = "Bonjour, "

	defaultName = "World"

	suffix = "!"
)

func Hello(name, language string) string {
	if name == "" {
		name = defaultName
	}

	return greetingPrefix(language) + name + suffix
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case romanian:
		prefix = romanianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
