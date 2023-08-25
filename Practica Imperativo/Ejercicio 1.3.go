//Ejercicio #3

package main

import (
	"fmt"
)

func rotateArray(arr []rune, n int, dir int) []rune {
	length := len(arr)

	// Determinar la dirección de rotación
	if dir == 0 { // Rotación a la izquierda
		for i := 0; i < n; i++ {
			first := arr[0]
			copy(arr[:], arr[1:])
			arr[length-1] = first
		}
	} else { // Rotación a la derecha
		for i := 0; i < n; i++ {
			last := arr[length-1]
			copy(arr[1:], arr[:length-1])
			arr[0] = last
		}
	}
	return arr
}

func main() {
	// Arreglo de ejemplo
	arr := []rune{'n', 'i', 'c', 'o', 'l', 'e'}
	fmt.Println("Arreglo original:", string(arr))

	// Rotación a la izquierda de 2 posiciones
	arr = rotateArray(arr, 2, 0)
	fmt.Println("Rotación a la izquierda:", string(arr))

	// Rotación a la derecha de 3 posiciones
	arr = rotateArray(arr, 3, 1)
	fmt.Println("Rotación a la derecha:", string(arr))
}
