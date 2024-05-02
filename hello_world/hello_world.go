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

	if language == romanian {
		return romanianHelloPrefix + name + suffix
	}

	if language == french {
		return  frenchHelloPrefix + name + suffix
	}

	return englishHelloPrefix + name + suffix
}

func main() {
	fmt.Println(Hello("", ""))
}
