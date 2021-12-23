package main

import "fmt"

func main() {

	var palabra string = "computadora"

	fmt.Println(len(palabra))

	for _, p := range palabra {

		fmt.Printf("%c \n", p)

	}

}
