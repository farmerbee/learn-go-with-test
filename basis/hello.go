package main

const spanish = "spanish"
const french = "french"
const englishPrefix = "hello "
const spanishPrefix = "Hola "
const frenchPrefix = "Bonjour "

func Hello(name, laguage string) string {
	if name == "" {
		name = "world"
	}
	return getPrefix(laguage) + name
}

func getPrefix(laguage string) (prefix string) {
	switch laguage {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}

	return
}
