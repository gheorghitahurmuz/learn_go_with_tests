package main

import "fmt"

const romanian = "Romanian"
const french = "French"
const englishHelloPrefix = "Hello, "
const romanianHelloPrefix = "Salut, "
const frenchHelloPrefix = "Bonjour, "
const defaultName = "World"
const suffix = "!"

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	prefix := englishHelloPrefix

	
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case romanian:
		prefix = romanianHelloPrefix
	}

	return prefix + name + suffix
}

func main() {
	fmt.Println(Hello("", ""))
}
