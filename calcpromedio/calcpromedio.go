package main

import (
	"errors"
	"fmt"
)

func calcPromedio(notas ...float64) (float64, error) {

	var resultado float64

	for _, a := range notas {
		if a <= 0 {

			fmt.Println("la nota ingresada", a, "no esta dentro de los valores permitidos")
			return 0, errors.New("debe ingresar una nota mayor o igual a 1")

		} else {

			resultado += a / float64(len(notas))
		}
	}
	return resultado, nil
}

func main() {

	resultado, error := calcPromedio(7, -6, -2, 5)

	if error != nil {

		fmt.Println(error)
	} else {

		fmt.Println("El promedio de las notas ingresadas es:", resultado)
	}

}
