package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero)

	factorial := 1

	for i := 1; i <= numero; i++ {

		factorial *= i

	}

	fmt.Println("Factorial:", factorial)

}