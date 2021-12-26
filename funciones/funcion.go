package main

import "fmt"

func operacionAritmetica(valor1, valor2 float64, operador string) float64 {

	switch operador {

	case "suma":
		return valor1 + valor2

	case "resta":
		return valor1 - valor2

	case "multiplicacion":
		return valor1 * valor2

	case "division":
		if valor2 != 0 {
			return valor1 / valor2

		}

	}
	return 0
}

func main() {

	fmt.Println(operacionAritmetica(5, 5, "suma"))

}
