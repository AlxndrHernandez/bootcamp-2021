package main

import "fmt"

func main() {

	var minutos int = 120
	var categoria string = "a"

	if categoria == "a" {

		horas := minutos / 60

		sueldo := float64(horas*3000) * 1.5

		fmt.Println("el total de tu sueldo es de :", sueldo)

	} else if categoria == "b" {

		horas := minutos / 60

		var sueldo float64 = float64(horas*1500) * 1.2

		fmt.Println("el total de tu sueldo es de :", sueldo)

	} else if categoria == "c" {

		horas := minutos / 60

		var sueldo float64 = float64(horas * 1000)

		fmt.Println("el total de tu sueldo es de :", sueldo)

	} else {
		fmt.Println("categoria invalida, ingrese categoria entre a, b y c")
	}

}
