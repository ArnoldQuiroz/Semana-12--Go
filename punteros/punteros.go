package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

func cumpleaniosCopia(p Persona) {
	p.Edad++
	fmt.Println("Dentro de la función (copia):", p.Edad)
}

func cumpleaniosPuntero(p *Persona) {
	p.Edad++
	fmt.Println("Dentro de la función (puntero):", p.Edad)
}

func nuevaPersona(nombre string, edad int) *Persona {
	return &Persona{
		Nombre: nombre,
		Edad:   edad,
	}
}

func intercambiar(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	
	var x int = 10
	var ptr *int = &x

	fmt.Println("Valor de x:", x)
	fmt.Println("Dirección de x:", ptr)
	fmt.Println("Valor apuntado por ptr:", *ptr)

	
	*ptr = 20
	fmt.Println("x después de *ptr = 20:", x)

	
	ptr2 := new(int)
	fmt.Println("Valor inicial con new:", *ptr2)
	*ptr2 = 100
	fmt.Println("Valor asignado:", *ptr2)
	ptr2 = nil 

	
	p1 := Persona{Nombre: "Ana", Edad: 25}

	
	cumpleaniosCopia(p1)
	fmt.Println("Ana después de cumpleaniosCopia:", p1.Edad) 

	
	cumpleaniosPuntero(&p1)
	fmt.Println("Ana después de cumpleaniosPuntero:", p1.Edad) 

	
	p2 := nuevaPersona("Carlos", 30)
	fmt.Printf("Persona creada: %s, %d años\n", p2.Nombre, p2.Edad)

	
	
	p2.Edad = 31 
	fmt.Println("Edad actualizada:", p2.Edad)

	
	a, b := 5, 10
	fmt.Printf("Antes: a=%d, b=%d\n", a, b)
	intercambiar(&a, &b)
	fmt.Printf("Después: a=%d, b=%d\n", a, b)

	
	p3 := &p1
	p4 := &p1
	fmt.Println("p3 y p4 apuntan al mismo lugar:", p3 == p4)

	p5 := &Persona{Nombre: "Luis", Edad: 20}
	fmt.Println("p3 y p5 apuntan al mismo lugar:", p3 == p5)
}
