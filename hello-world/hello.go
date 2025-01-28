package hello

import (
	"fmt"
)

const (
	portuguese = "Portuguese"
	french     = "French"

	englishHelloPrefix    = "Hello, "
	portugueseHelloPrefix = "Olá, "
	frenchHelloPrefix     = "Bonjour, "

	exclamationMark = "!"
)

func nameIsEmpty(name string) (answer bool) {
	if name == "" {
		answer = true
	} else {
		answer = false
	}
	return
}

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portugueseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if nameIsEmpty(name) {
		name = "World"
	}

	return getGreetingPrefix(language) + name + exclamationMark
}

func main() {
	fmt.Println(Hello("Lucas", "Portuguese"))
}
