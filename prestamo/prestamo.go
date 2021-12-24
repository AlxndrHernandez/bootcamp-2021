package main

import "fmt"

func main() {

	var edad int = 25
	var empleado bool = true
	var antiguedadLaboral = 0
	var sueldo int = 200000

	switch {
	case edad < 22:
		fmt.Println("tu edad es menor a 22 años, no puedes acceder a un credito ")
		fallthrough

	case !empleado:
		fmt.Println("Estas desempleado, por ende no puedes acceder a un credito")
		fallthrough

	case antiguedadLaboral < 1:
		fmt.Println("tu antiguedad laboral es menor a un año no puedes acceder a un credito")
		fallthrough

	case sueldo < 100000:
		fmt.Println("al tener un suedo menor a 100.000 se te cobrará interes")

	default:
		fmt.Println("cumples con los requisitos")

	}

}
