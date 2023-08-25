package main

import (
	"fmt"
)

// estructura calzado
type calzado struct {
	marca  string
	precio int32
	talla  int32
	stock  int32
}

type zapatos []calzado

var lista zapatos

const tallaMenor = 34
const tallaMayor = 44

func (l *zapatos) agregarZapato(mar string, prec int32, num int32, stock int32) {
	idx := buscarMarcaTalla(*l, mar, num) // Usamos la función auxiliar
	if idx == -1 {
		*l = append(*l, calzado{mar, prec, num, stock})
	} else {
		fmt.Println("Zapato ya existe en el inventario")
	}
}

func buscarMarcaTalla(list zapatos, marca string, talla int32) int {
	for i, p := range list {
		if p.marca == marca && p.talla == talla {
			return i
		}
	}
	return -1
}

// vende zapatos y los elimina del stock
func (l *zapatos) venderZapato(mar string, num int32) {
	idx := buscarMarcaTalla(*l, mar, num)
	if idx == -1 {
		fmt.Println("Zapato no encontrado en el inventario")
		return
	}

	if (*l)[idx].stock <= 0 {
		fmt.Println("No hay stock disponible para este zapato")
		return
	}

	(*l)[idx].stock--
	fmt.Println("Venta realizada con éxito")
}

func main() {
	lista.agregarZapato("Nike", 35000, 36, 5)
	lista.agregarZapato("Adidas3", 32500, 40, 10)
	lista.agregarZapato("Coral", 24300, 38, 3)
	lista.agregarZapato("NikeAir", 18900, 42, 29)
	lista.agregarZapato("Adidas2", 31800, 40, 11)
	lista.agregarZapato("Coral2", 18000, 39, 8)
	lista.agregarZapato("Nike", 35000, 37, 5)
	lista.agregarZapato("Adidas", 32500, 35, 10)
	lista.agregarZapato("Coral3", 25300, 43, 3)
	lista.agregarZapato("Nike", 46700, 36, 5)
	lista.agregarZapato("Adidas1", 37500, 35, 10)
	lista.agregarZapato("Coral4", 29300, 38, 3)

	// Realizar ventas de ejemplo
	lista.venderZapato("Nike", 36)
	lista.venderZapato("Adidas", 40)
	lista.venderZapato("Coral", 38)

	// Intentar vender un zapato sin stock
	lista.venderZapato("Nike", 37)
}
