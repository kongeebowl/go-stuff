package main

const (
	spanish = "Spanish"
	french  = "French"

	englishHello = "Hello, "
	spanishHello = "Hola, "
	frenchHello  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greeting(language) + name
}

func greeting(language string) (prefix string) {
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
