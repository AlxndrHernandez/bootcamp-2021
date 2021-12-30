package main

import (
	"fmt"
)

type Matrix struct {
	valores    []float64
	alto       int
	ancho      int
	cuadratica bool
	maxvalue   float64
}

func (m *Matrix) set(numeros ...float64) {

	m.valores = numeros

}

func (m Matrix) print() {

	if len(m.valores) == 0 {
		fmt.Println("La matriz está vacía")
	}
	for fila := 0; fila <= m.alto; fila++ {
		fmt.Println(m.valores[fila*m.ancho : fila*m.ancho+m.ancho])
	}

	if m.alto == m.ancho {

		fmt.Println("La matriz es cuadratica")
		m.cuadratica = true
	} else {

		fmt.Println("La matriz no es cuadratica")
		m.cuadratica = false
	}
	numeroMayor := m.valores[0]
	for _, v := range m.valores {
		if v > numeroMayor {
			numeroMayor = v
		}

	}
	fmt.Printf("EL numero mayor es: %v", numeroMayor)
}

func main() {

	m := Matrix{

		alto:       3,
		ancho:      3,
		cuadratica: true,
		maxvalue:   0,
	}
	fmt.Println(m.valores)
	m.set(1, 2, 3, 4, 5, 6, 7, 8, 9, 99, 11, 12)
	fmt.Println(m.valores)
	m.print()

}
