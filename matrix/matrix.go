package main

import "fmt"

type Matrix struct {
	valores []float64
	alto    float64
	ancho   float64
}

func (m Matrix) set() {

	fmt.Printf("\n Valores: %v\n Alto: %v\n Ancho: %v\n", m.valores, m.alto, m.ancho)

}

func main() {

	matrix := Matrix{

		valores: []float64{2, 3, 4, 5},
		alto:    4,
		ancho:   4,
	}

	matrix.set()

}
