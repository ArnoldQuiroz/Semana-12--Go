package main

import "fmt"

type Estudiante struct {
	ID       int
	Nombre   string
	Carrera  string
	Promedio float64
	Activo   bool
}

type Curso struct {
	ID     int
	Nombre string
	Creditos int
}

type Universidad struct {
	Nombre     string
	Direccion  string
	Estudiantes []Estudiante
	Cursos      []Curso
}

func main() {
	
	var e1 Estudiante
	fmt.Println("Estudiante vacío:", e1)

	
	e2 := Estudiante{
		ID:       1,
		Nombre:   "Jose",
		Carrera:  "Ingenieria de Sistemas",
		Promedio: 18.5,
		Activo:   true,
	}
	fmt.Println("Estudiante:", e2)

	
	e3 := Estudiante{
		Nombre:  "Maria",
		Carrera: "Medicina",
	}
	fmt.Println("Estudiante parcial:", e3)

	
	e4 := &Estudiante{
		ID:     2,
		Nombre: "Pedro",
	}
	fmt.Println("Puntero a estudiante:", e4)

	
	e2.Promedio = 19.0
	fmt.Println("Promedio actualizado:", e2.Promedio)

	
	profesor := struct {
		Nombre   string
		Materia  string
		Antiguedad int
	}{
		Nombre:   "Dr. García",
		Materia:  "Algoritmos",
		Antiguedad: 10,
	}
	fmt.Println("Profesor:", profesor)

	
	uni := Universidad{
		Nombre:    "USS",
		Direccion: "Av. Principal 123",
		Estudiantes: []Estudiante{e2, e3},
		Cursos: []Curso{
			{ID: 1, Nombre: "Programación", Creditos: 4},
			{ID: 2, Nombre: "Matemáticas", Creditos: 3},
		},
	}
	fmt.Printf("Universidad: %s con %d estudiantes\n", uni.Nombre, len(uni.Estudiantes))
}
