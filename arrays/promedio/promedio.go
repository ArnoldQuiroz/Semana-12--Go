package main

import "fmt"

func main() {

	notas := [5]float64{18, 17, 15, 20, 16}

	suma := 0.0

	for i := 0; i < len(notas); i++ {

		suma += notas[i]

	}

	promedio := suma / float64(len(notas))

	fmt.Println("Promedio:", promedio)

}