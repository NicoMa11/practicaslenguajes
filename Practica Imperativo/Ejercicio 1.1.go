//Ejercicio #1

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Soy Nicole y me gusta bailar folklore"
	fmt.Println(text)

	// Contar caracteres
	numChars := len(text)

	// Contar palabras
	words := splitBySpace(text)
	numWords := len(words)

	// Contar líneas
	numLines := strings.Count(text, "\n") + 1

	// Mostrar resultados
	fmt.Printf("Número de caracteres: %d\n", numChars)
	fmt.Printf("Número de palabras: %d\n", numWords)
	fmt.Printf("Número de líneas: %d\n", numLines)
}

func splitBySpace(text string) []string {
	// Dividir el texto por espacios en blanco
	return strings.Fields(text)
}
