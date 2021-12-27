package main

import "fmt"

type Estudiantes struct {
	nombre   string
	apellido string
	dni      string
	fecha    string
}

func (e Estudiantes) Detalle() {

	fmt.Println("\nDetalle del estudiantes")
	fmt.Printf("\nNombre: %s\nApellido: %s\nDni: %s\nFecha: %s\n", e.nombre, e.apellido, e.dni, e.fecha)

}

func main() {

	estudiante := Estudiantes{

		nombre:   "Alexander",
		apellido: "Hernández",
		dni:      "17781063-0",
		fecha:    "27/03/1991",
	}

	estudiante2 := Estudiantes{

		nombre:   "Francisco",
		apellido: "Sandoval",
		dni:      "19984887-5",
		fecha:    "08/05/1998",
	}

	// fmt.Println(estudiante)
	// fmt.Println(estudiante2)
	estudiante.Detalle()
	estudiante2.Detalle()

}
