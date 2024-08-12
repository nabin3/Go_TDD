package helloWorld

import (
	"errors"
	"strings"
)

const (
	spanish = "spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

var errArgumentNumber = errors.New("hello func can't take more than two arguments")

/*
Hello func takes take two arguments, 1st is for language and 2nd is for name.
It can be passed no argument at all or only language.
Now if someOne pass only name of someone it will return only "Hello, World".
*/
func Hello(params ...string) (string, error) {
	paramsNumber := len(params)

	if paramsNumber > 2 {
		return "", errArgumentNumber
	}

	var language string
	if paramsNumber == 0 {
		language = "english"
	} else {
		language = params[0]
	}

	var name string
	if paramsNumber == 0 || paramsNumber == 1 || params[1] == "" {
		name = "World"
	} else {
		name = params[1]
	}

	return greetingPrefix(language) + name, nil
}

func greetingPrefix(language string) string {
	var helloPrefix string

	switch strings.ToLower(language) {
	case spanish:
		helloPrefix = spanishHelloPrefix
	default:
		helloPrefix = englishHelloPrefix
	}

	return helloPrefix
}
