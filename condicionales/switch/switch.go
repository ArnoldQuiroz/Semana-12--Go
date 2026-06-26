package main

import "fmt"

func main() {

	dia := 3

	switch dia {

	case 1:
		fmt.Println("Lunes")

	case 2:
		fmt.Println("Martes")

	case 3:
		fmt.Println("Miércoles")

	case 4:
		fmt.Println("Jueves")

	case 5:
		fmt.Println("Viernes")

	case 6:
		fmt.Println("Sábado")

	case 7:
		fmt.Println("Domingo")

	default:
		fmt.Println("Día inválido")
	}

}