package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

const spanishLang = "Spanish"
const englishLang = "English"
const frenchLang = "French"

// [01]
/* func Hello(name string) string {
  return fmt.Sprintf("%s%s", englishHelloPrefix, name)
} */

// [02]
/* func Hello(name string) string {
  const defaultName = "world"

  if name == "" {
    name = defaultName
  }

  return fmt.Sprintf("%s%s", englishHelloPrefix, name)
} */

// [03]
/* func Hello(name, language string) string {
	const defaultName = "world"

	if name == "" {
		name = defaultName
	}

	switch language {
	case spanishLang:
		return fmt.Sprintf("%s%s", spanishHelloPrefix, name)

	case frenchLang:
		return fmt.Sprintf("%s%s", frenchHelloPrefix, name)

	default:
		return fmt.Sprintf("%s%s", englishHelloPrefix, name)
	}
} */

// [03]
func Hello(name, language string) string {
	const defaultName = "world"

	if name == "" {
		name = defaultName
	}

  prefix := grettingPrefix(language)

  return fmt.Sprintf("%s%s", prefix, name)
}

func grettingPrefix(language string) string {
	switch language {
	case spanishLang:
		return spanishHelloPrefix

	case frenchLang:
		return frenchHelloPrefix

	default:
		return englishHelloPrefix
	}

}

func mainHello() {
	fmt.Println(Hello("Jams", "Spanish"))
}
