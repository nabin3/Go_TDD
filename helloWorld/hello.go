package helloWorld

const (
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

func Hello(language, name string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	var helloPrefix string
	switch language {
	case spanish:
		helloPrefix = spanishHelloPrefix
	default:
		helloPrefix = englishHelloPrefix
	}

	return helloPrefix
}
