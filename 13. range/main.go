package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// Palindromo usando un string como arreglo y parsearlo a string
// porque su resultado es caracter (byte)
func palindromo(text string) {
	var textReverse string
	lower := strings.ToLower(text)
	for i:= len(lower) - 1; i >= 0; i-- {
		textReverse += string(lower[i])
	}
	if lower == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	
	// Uso de RANGE
	slice := []string{"Hola", "mundo!", "desde", "un", "Slice"}

	for index, value := range slice {
		fmt.Println(index, value)
	}

	// Solo el segundo elemento
	for _, value := range slice {
		fmt.Println(value)
	}

	slice[0] = "konichiwa"
	fmt.Println(slice)
	
	// Palindromo
	
	fmt.Println(isPalindrome("A man a plan a canal Panama"))
	palindromo("Ama")
	fmt.Println(strings.ToLower("Ama"))
}