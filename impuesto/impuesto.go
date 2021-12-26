package main

import "fmt"

func impuesto(valor1 float64) float64 {

	if valor1 > 150000 {

		return (valor1 - (valor1 * 0.27))

	} else if valor1 > 50000 {

		return (valor1 - (valor1 * 0.17))

	} else {

		return (valor1)
	}

}

func main() {

	fmt.Println(impuesto(150000))

}
