package main

import "fmt"

type Language string

const helloPrefix = "Hello"

const (
	LanguageEnglish   Language = "en"
	LanguageSpanish   Language = "sp"
	LanguageIndonesia Language = "id"
	LanguageArabic    Language = "ar"
)

var greetingPrefixes = map[Language]string{
	LanguageEnglish:   "Hello",
	LanguageSpanish:   "Hola",
	LanguageIndonesia: "Halo",
	LanguageArabic:    "Ahlan",
}

func (l Language) IsValid() bool {
	switch l {
	case LanguageEnglish, LanguageSpanish, LanguageIndonesia, LanguageArabic:
		return true
	default:
		return false
	}
}

func (l Language) greetingPrefix() string {
	greeting, ok := greetingPrefixes[l]
	if !ok {
		return greetingPrefixes[LanguageEnglish]
	}

	return greeting
}

func Hello(who string, language string) string {
	if who == "" {
		who = "World"
	}

	lang := Language(language)

	return fmt.Sprintf("%s, %s", lang.greetingPrefix(), who)
}

func main() {
	fmt.Println(Hello("Fulan", "wo"))
}
