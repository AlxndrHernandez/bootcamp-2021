package main

import "fmt"

func main() {

	precio := 15000
	descuento := 30
	total := (precio * (100 - descuento)) / 100

	fmt.Println("el precio total del producto aplicando el descuento, es de:", total)

}

//hola
