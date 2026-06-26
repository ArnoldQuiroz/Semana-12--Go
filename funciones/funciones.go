package main

import (
	"fmt"
	"strings"
)

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("división por cero no permitida")
	}
	return a / b, nil
}

func analizarTexto(texto string) (palabras int, caracteres int, lineas int) {
	caracteres = len(texto)
	lineas = strings.Count(texto, "\n") + 1
	palabras = len(strings.Fields(texto))
	return
}

func sumar(numeros ...int) int {
	total := 0
	for _, n := range numeros {
		total += n
	}
	return total
}

func registrarLog(nivel string, mensajes ...string) {
	fmt.Printf("[%s]", nivel)
	for _, m := range mensajes {
		fmt.Printf(" %s", m)
	}
	fmt.Println()
}

type Operacion func(int, int) int

func aplicar(a, b int, op Operacion) int {
	return op(a, b)
}

func crearMultiplicador(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func crearContador() func() int {
	cuenta := 0
	return func() int {
		cuenta++
		return cuenta
	}
}

func crearAcumulador(inicial float64) func(float64) float64 {
	total := inicial
	return func(valor float64) float64 {
		total += valor
		return total
	}
}

func filtrar(nums []int, condicion func(int) bool) []int {
	resultado := []int{}
	for _, n := range nums {
		if condicion(n) {
			resultado = append(resultado, n)
		}
	}
	return resultado
}

func mapear(nums []int, transformar func(int) int) []int {
	resultado := make([]int, len(nums))
	for i, n := range nums {
		resultado[i] = transformar(n)
	}
	return resultado
}

func main() {
	
	fmt.Println("=== Múltiples retornos ===")
	resultado, err := dividir(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", resultado)
	}

	_, err = dividir(5, 0)
	if err != nil {
		fmt.Println("Error esperado:", err)
	}

	
	palabras, chars, lineas := analizarTexto("Hola mundo\ncomo estás\nbien gracias")
	fmt.Printf("Palabras: %d, Caracteres: %d, Líneas: %d\n", palabras, chars, lineas)

	
	fmt.Println("\n=== Funciones variádicas ===")
	fmt.Println("Suma:", sumar(1, 2, 3, 4, 5))
	fmt.Println("Suma:", sumar(10, 20))

	nums := []int{1, 2, 3}
	fmt.Println("Suma desde slice:", sumar(nums...))

	registrarLog("INFO", "sistema iniciado", "conexión establecida")
	registrarLog("ERROR", "fallo crítico")

	
	fmt.Println("\n=== Funciones como parámetros ===")
	suma := func(a, b int) int { return a + b }
	resta := func(a, b int) int { return a - b }
	multi := func(a, b int) int { return a * b }

	fmt.Println("aplicar suma(3,4):", aplicar(3, 4, suma))
	fmt.Println("aplicar resta(10,3):", aplicar(10, 3, resta))
	fmt.Println("aplicar multi(5,6):", aplicar(5, 6, multi))

	
	fmt.Println("\n=== Closures ===")
	doble := crearMultiplicador(2)
	triple := crearMultiplicador(3)
	fmt.Println("doble(5):", doble(5))
	fmt.Println("triple(5):", triple(5))

	
	contador1 := crearContador()
	contador2 := crearContador()
	fmt.Println("contador1:", contador1(), contador1(), contador1())
	fmt.Println("contador2:", contador2(), contador2())

	
	caja := crearAcumulador(0)
	fmt.Println("caja:", caja(100), caja(50), caja(25))

	
	fmt.Println("\n=== Orden superior ===")
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	pares := filtrar(numeros, func(n int) bool { return n%2 == 0 })
	fmt.Println("Pares:", pares)

	mayoresDe5 := filtrar(numeros, func(n int) bool { return n > 5 })
	fmt.Println("Mayores de 5:", mayoresDe5)

	cuadrados := mapear(numeros, func(n int) int { return n * n })
	fmt.Println("Cuadrados:", cuadrados)
}
