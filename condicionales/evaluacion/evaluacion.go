package main

import "fmt"

func main() {

	var nota int

	fmt.Print("Ingrese su nota: ")
	fmt.Scan(&nota)

	if nota >= 18 {
		fmt.Println("Excelente")
	} else if nota >= 14 {
		fmt.Println("Aprobado")
	} else {
		fmt.Println("Desaprobado")
	}

}