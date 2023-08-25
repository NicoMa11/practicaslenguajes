//Ejercicio #2

package main

import "fmt"

// Imprime una línea con n asteriscos
func imprimirLineaAsteriscos(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("* ")
	}
	fmt.Println()
}

// Imprimir una línea con espacios en blanco
func imprimirLineaEspacios(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("  ")
	}
}

// Imprime figura en forma de pirámide
func imprimirFigura(numElement int) {
	if numElement%2 == 0 || numElement < 0 {
		fmt.Println("Ingrese un número impar positivo")
		return
	}

	numEspacios := (numElement - 1) / 2
	numAsteriscos := 1

	//Ciclo para imprimir los asteriscos y espacios correspondientes
	for i := 0; i < numElement; i++ {
		imprimirLineaEspacios(numEspacios)
		imprimirLineaAsteriscos(numAsteriscos)

		if i < numElement/2 {
			numEspacios--
			numAsteriscos += 2
		} else {
			numEspacios++
			numAsteriscos -= 2
		}
	}
}

func main() {
	imprimirFigura(7)
}
