package main

import "fmt"

func main() {

	aplicarDescuento(23.4, 12.3)

}

func aplicarDescuento(precio float64, descuento float64) {

	var valorResta float64
	valorResta = calcularDescuento(precio, descuento)

	fmt.Println(valorResta)

}

func calcularDescuento(precio float64, descuento float64) float64 {

	var val float64
	val = descuento * precio / 100

	return val

}
