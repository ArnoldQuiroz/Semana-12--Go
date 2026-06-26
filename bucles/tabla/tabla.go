package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero)

	for i := 1; i <= 12; i++ {

		fmt.Printf("%d x %d = %d\n",
			numero,
			i,
			numero*i)

	}

}