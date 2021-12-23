package main

import "fmt"

func main() {

	var palabra string = "computadora"
	pointer := &palabra
	fmt.Printf("la direccion de memoria es %p, %p y su contenido es %s \n", pointer, &palabra, *pointer)
	fmt.Println(len(palabra))
	fmt.Printf("%s \n", palabra)
	for _, p := range palabra {

		fmt.Printf("%c \n", p)

	}
	//hola
}
