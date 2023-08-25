package main

import (
	"fmt"
	"sort"
)

// estructura producto
type producto struct {
	nombre   string
	cantidad int
	precio   int
}

type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10

// se agregan productos con datos quemados a la lista productos
func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	if prod, err := l.buscarProducto(nombre); err == 0 {
		prod.cantidad += cantidad
		if prod.precio != precio {
			prod.precio = precio
		}
	} else {
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}
}

// agrega productos con datos quemados
func (l *listaProductos) agregarProductos(productos ...producto) {
	for _, p := range productos {
		l.agregarProducto(p.nombre, p.cantidad, p.precio)
	}
}

// Busca los productos, se utiliza en otras funciones para modificar los productos o el inventario
func (l *listaProductos) buscarProducto(nombre string) (*producto, int) {
	for i, p := range *l {
		if p.nombre == nombre {
			return &(*l)[i], 0 // Producto encontrado y sin error
		}
	}
	return nil, -1 // Producto no encontrado
}

// busca el producto y lo elimina del inventario
func (l *listaProductos) venderProducto(nombre string, cantidad int) {
	if prod, err := l.buscarProducto(nombre); err == 0 {
		if prod.cantidad >= cantidad {
			prod.cantidad -= cantidad
			if prod.cantidad == 0 {
				l.eliminarProducto(nombre)
				fmt.Printf("Producto %s agotado y eliminado del inventario\n", nombre)
			}
		} else {
			fmt.Println("No hay suficiente stock para vender")
		}
	} else {
		fmt.Println("Producto no encontrado en el inventario")
	}
}

func (l *listaProductos) eliminarProducto(nombre string) {
	for i, p := range *l {
		if p.nombre == nombre {
			*l = append((*l)[:i], (*l)[i+1:]...)
			break
		}
	}
}

func llenarDatos() {
	lProductos.agregarProductos(
		producto{"arroz", 15, 2500},
		producto{"frijoles", 4, 2000},
		producto{"leche", 8, 1200},
		producto{"café", 12, 4500},
	)
}

// crea una lista de productos con cantidades mínimas en el inventario
func (l *listaProductos) listarProductosMinimos() listaProductos {
	newSlice := make(listaProductos, 0)

	for _, p := range *l {
		if p.cantidad <= existenciaMinima { //compara la cantidad del producto con la existencia minima = 10
			newSlice = append(newSlice, p)
		}
	}
	return newSlice
}

// Se  completa la cantidad del producto, para llegue a la cantiad mínima
func (l *listaProductos) aumentarInventarioDeMinimos(listaMinimos listaProductos) {
	for _, p := range listaMinimos {
		cantidadNecesaria := existenciaMinima - p.cantidad
		if cantidadNecesaria > 0 {
			l.agregarProducto(p.nombre, cantidadNecesaria, p.precio)
			fmt.Printf("Se agregaron %d unidades de %s al inventario\n", cantidadNecesaria, p.nombre)
		}
	}
}

// ordena los productos por nombre, cantidad y precio
func (l *listaProductos) ordenarPorCampo(campo string) {
	switch campo {
	case "nombre":
		sort.SliceStable(*l, func(i, j int) bool {
			return (*l)[i].nombre < (*l)[j].nombre
		})
	case "cantidad":
		sort.SliceStable(*l, func(i, j int) bool {
			return (*l)[i].cantidad < (*l)[j].cantidad
		})
	case "precio":
		sort.SliceStable(*l, func(i, j int) bool {
			return (*l)[i].precio < (*l)[j].precio
		})
	default:
		fmt.Println("Campo de ordenamiento no válido")
	}
}

func main() {
	llenarDatos()
	fmt.Println("Inventario inicial:")
	fmt.Println(lProductos)

	lProductos.venderProducto("arroz", 2)
	lProductos.venderProducto("frijoles", 3)
	lProductos.venderProducto("café", 5)
	fmt.Println("Inventario después de ventas:")
	fmt.Println(lProductos)

	pMinimos := lProductos.listarProductosMinimos()
	fmt.Println("Productos con existencia mínima:")
	fmt.Println(pMinimos)

	lProductos.aumentarInventarioDeMinimos(pMinimos)
	fmt.Println("Inventario después de aumentar existencias:")
	fmt.Println(lProductos)

	fmt.Println("Ordenando por nombre:")
	lProductos.ordenarPorCampo("nombre")
	fmt.Println(lProductos)

	fmt.Println("Ordenando por cantidad:")
	lProductos.ordenarPorCampo("cantidad")
	fmt.Println(lProductos)

	fmt.Println("Ordenando por precio:")
	lProductos.ordenarPorCampo("precio")
	fmt.Println(lProductos)
}
