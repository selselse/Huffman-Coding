package main

import (
	"fmt"
)

func harfdeneme() {
	str := "こんにちは" // Japanese word for "hello"

	// Accessing Unicode code point of the first character in the string
	firstRune := rune(str[0])
	fmt.Printf("Unicode code point of the first character: %v\n", firstRune)

	// Accessing Unicode code point of the second character in the string
	secondRune := rune(str[1])
	fmt.Printf("Unicode code point of the second character: %v\n", secondRune)
}
