package main

import "fmt"

func main() {

	numeros := [5]int{10, 20, 30, 40, 50}

	for i := 0; i < len(numeros); i++ {

		fmt.Println(numeros[i])

	}

}