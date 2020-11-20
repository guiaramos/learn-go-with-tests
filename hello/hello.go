package main

import "fmt"

const (
	portuguese            = "Portuguese"
	french                = "French"
	englishHelloPrefix    = "Hello, "
	portugueseHelloPrefix = "Ola, "
	frenchHelloPrefix     = "Bonjour, "
)

// Hello returns hello to world
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", "english"))
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
