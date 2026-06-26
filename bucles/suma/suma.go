package main

import "fmt"

func main() {

	suma := 0

	for i := 1; i <= 100; i++ {

		suma += i

	}

	fmt.Println("Resultado:", suma)

}