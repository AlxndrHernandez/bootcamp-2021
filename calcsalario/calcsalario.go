package main

import "fmt"

func main() {

	var minutos int = 2600
	var categoria string = "c"

	if categoria == "a" {

		var horas float64 = float64(minutos) / 60
		fmt.Println(horas)

		sueldo := float64(horas*3000) * 1.5

		fmt.Println("el total de tu sueldo es de :", sueldo)

	} else if categoria == "b" {

		var horas float64 = float64(minutos) / 60

		var sueldo float64 = float64(horas*1500) * 1.2

		fmt.Println("el total de tu sueldo es de :", sueldo)

	} else if categoria == "c" {

		var horas float64 = float64(minutos) / 60

		sueldo := horas * 1000

		fmt.Println("el total de tu sueldo es de :", sueldo)

	} else {
		fmt.Println("categoria invalida, ingrese categoria entre a, b y c")
	}

}
