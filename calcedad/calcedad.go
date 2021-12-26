package main

import "fmt"

func main() {

	var employees = map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Darió":    44,
		"Pedro":    30}

	fmt.Println("La edad de benjamin es:", employees["Benjamin"])

	contador := 0
	for _, value := range employees {
		if value >= 21 {

			contador++

		}
	}
	fmt.Println("En la empresa hay:", contador, "empleados que tienen 21 años o mas")

	fmt.Println("Se agrega al empleado Federico")
	employees["Federico"] = 25
	fmt.Println(employees)
	fmt.Println("Se elimina al empleado Pedro")
	delete(employees, "Pedro")
	fmt.Println(employees)

}
