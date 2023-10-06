package main

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "

	spanish = "Spanish"
	french  = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := gettingPrefix(language)

	return prefix + name
}

func gettingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
