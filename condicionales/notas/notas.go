package main

import "fmt"

func main() {

	nota := 17

	if nota >= 18 {
		fmt.Println("Excelente")
	} else if nota >= 14 {
		fmt.Println("Aprobado")
	} else {
		fmt.Println("Desaprobado")
	}

}