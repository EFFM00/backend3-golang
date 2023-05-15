package main

import (
	"errors"
	"fmt"
	"log"
)

const (
	perro       = "perro"
	grPerro     = 10000
	gato        = "gato"
	grGato      = 5000
	tarantula   = "tarantula"
	grTarantula = 150
	hamster     = "hamster"
	grHamster   = 250
)

func main() {

	alimentarPerros, err := animal(perro)

	if err != nil {
		log.Fatal(err)
	}

	alimentarGatos, err := animal(gato)

	if err != nil {
		log.Fatal(err)
	}

	alimentarTarantulas, err := animal(tarantula)

	if err != nil {
		log.Fatal(err)
	}

	alimentarHamsters, err := animal(hamster)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("La comida para 5 perros es %v gr. \n", alimentarPerros(5))
	fmt.Printf("La comida para 2 gatos es %v gr. \n", alimentarGatos(2))
	fmt.Printf("La comida para 6 tarántulas es %v gr. \n", alimentarTarantulas(6))
	fmt.Printf("La comida para 3 hámsters es %v gr. \n", alimentarHamsters(3))
}

func animal(animal string) (func(int) float64, error) {

	switch animal {

	case perro:

		return alimentarPerro, nil

	case gato:

		return alimentarGato, nil

	case tarantula:

		return alimentarTarantula, nil

	case hamster:

		return alimentarHamster, nil

	default:

		return nil, errors.New("Animal no encontrado")

	}

}

func alimentarPerro(cantidadPerros int) float64 {

	return float64(cantidadPerros) * grPerro

}

func alimentarGato(cantidadGatos int) float64 {

	return float64(cantidadGatos) * grGato

}

func alimentarTarantula(cantidadTarantulas int) float64 {

	return float64(cantidadTarantulas) * grTarantula

}

func alimentarHamster(cantidadHamsters int) float64 {

	return float64(cantidadHamsters) * grHamster

}
