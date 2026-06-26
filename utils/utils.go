package main

import (
	"fmt"
	"sort"
	"strings"
)

type Libro struct {
	Titulo   string
	Autor    string
	Paginas  int
	Genero   string
	Calificacion float64
}

type FiltroLibro func(Libro) bool
type ComparadorLibro func(Libro, Libro) bool
type TransformadorLibro func(Libro) string

func FiltrarLibros(libros []Libro, filtro FiltroLibro) []Libro {
	resultado := make([]Libro, 0)
	for _, libro := range libros {
		if filtro(libro) {
			resultado = append(resultado, libro)
		}
	}
	return resultado
}

func MapearLibros(libros []Libro, transformador TransformadorLibro) []string {
	resultado := make([]string, len(libros))
	for i, libro := range libros {
		resultado[i] = transformador(libro)
	}
	return resultado
}

func OrdenarLibros(libros []Libro, comparador ComparadorLibro) []Libro {
	copia := make([]Libro, len(libros))
	copy(copia, libros)
	sort.Slice(copia, func(i, j int) bool {
		return comparador(copia[i], copia[j])
	})
	return copia
}

func ReducirLibros(libros []Libro, inicial float64, acumulador func(float64, Libro) float64) float64 {
	resultado := inicial
	for _, libro := range libros {
		resultado = acumulador(resultado, libro)
	}
	return resultado
}

func CrearFiltroAutor(autor string) FiltroLibro {
	return func(libro Libro) bool {
		return strings.EqualFold(libro.Autor, autor)
	}
}

func CrearFiltroPaginas(min, max int) FiltroLibro {
	return func(libro Libro) bool {
		return libro.Paginas >= min && libro.Paginas <= max
	}
}

func CrearFiltroGenero(genero string) FiltroLibro {
	return func(libro Libro) bool {
		return strings.EqualFold(libro.Genero, genero)
	}
}

func CrearFiltroCalificacion(minCalificacion float64) FiltroLibro {
	return func(libro Libro) bool {
		return libro.Calificacion >= minCalificacion
	}
}

func CombinarFiltros(filtros ...FiltroLibro) FiltroLibro {
	return func(libro Libro) bool {
		for _, filtro := range filtros {
			if !filtro(libro) {
				return false
			}
		}
		return true
	}
}

func CombinarFiltrosOR(filtros ...FiltroLibro) FiltroLibro {
	return func(libro Libro) bool {
		for _, filtro := range filtros {
			if filtro(libro) {
				return true
			}
		}
		return false
	}
}

func createCounter() func() int {
	contador := 0
	return func() int {
		contador++
		return contador
	}
}

func createCounterConReset() (func() int, func()) {
	contador := 0
	incrementar := func() int {
		contador++
		return contador
	}
	resetear := func() {
		contador = 0
	}
	return incrementar, resetear
}

func CrearAcumuladorCalificaciones() func(float64) float64 {
	suma := 0.0
	cantidad := 0
	return func(calificacion float64) float64 {
		suma += calificacion
		cantidad++
		if cantidad == 0 {
			return 0
		}
		return suma / float64(cantidad)
	}
}

func main() {
	fmt.Println("FUNCIONES DE PRIMERA CLASE Y CLOSURES")
	fmt.Println(strings.Repeat("=", 50))

	
	biblioteca := []Libro{
		{"El Quijote", "Cervantes", 863, "Clásico", 4.8},
		{"Cien Años de Soledad", "García Márquez", 471, "Realismo Mágico", 4.9},
		{"Go Programming", "Donovan", 380, "Tecnología", 4.5},
		{"Clean Code", "Martin", 464, "Tecnología", 4.7},
		{"The Hobbit", "Tolkien", 310, "Fantasía", 4.6},
		{"1984", "Orwell", 328, "Distopía", 4.7},
		{"Dune", "Herbert", 896, "Ciencia Ficción", 4.8},
		{"Effective Go", "Donovan", 214, "Tecnología", 4.3},
	}

	
	fmt.Println("\n1. FILTROS BÁSICOS:")
	fmt.Println(strings.Repeat("-", 40))

	librosGrandes := FiltrarLibros(biblioteca, func(l Libro) bool {
		return l.Paginas > 400
	})
	fmt.Println("Libros con más de 400 páginas:")
	for _, l := range librosGrandes {
		fmt.Printf("  - %s (%d páginas)\n", l.Titulo, l.Paginas)
	}

	
	fmt.Println("\n2. CLOSURES: FÁBRICAS DE FILTROS:")
	fmt.Println(strings.Repeat("-", 40))

	filtroTecnologia := CrearFiltroGenero("tecnología")
	librosTec := FiltrarLibros(biblioteca, filtroTecnologia)
	fmt.Println("Libros de Tecnología:")
	for _, l := range librosTec {
		fmt.Printf("  - %s por %s\n", l.Titulo, l.Autor)
	}

	filtroDonovan := CrearFiltroAutor("donovan")
	librosDonovan := FiltrarLibros(biblioteca, filtroDonovan)
	fmt.Println("Libros de Donovan:")
	for _, l := range librosDonovan {
		fmt.Printf("  - %s (%d págs)\n", l.Titulo, l.Paginas)
	}

	
	fmt.Println("\n3. COMBINACIÓN DE FILTROS:")
	fmt.Println(strings.Repeat("-", 40))

	filtroTecYBuenoCalificado := CombinarFiltros(
		CrearFiltroGenero("tecnología"),
		CrearFiltroCalificacion(4.5),
	)
	librosTecBuenos := FiltrarLibros(biblioteca, filtroTecYBuenoCalificado)
	fmt.Println("Libros de Tecnología con calificación >= 4.5:")
	for _, l := range librosTecBuenos {
		fmt.Printf("  - %s (%.1f)\n", l.Titulo, l.Calificacion)
	}

	filtroMediano := CombinarFiltros(
		CrearFiltroPaginas(300, 500),
		CrearFiltroCalificacion(4.5),
	)
	librosMedianoBuenos := FiltrarLibros(biblioteca, filtroMediano)
	fmt.Println("\nLibros entre 300-500 páginas con calificación >= 4.5:")
	for _, l := range librosMedianoBuenos {
		fmt.Printf("  - %s (%d págs, %.1f)\n", l.Titulo, l.Paginas, l.Calificacion)
	}

	
	fmt.Println("\n4. TRANSFORMACIONES CON MAP:")
	fmt.Println(strings.Repeat("-", 40))

	resumenLibros := MapearLibros(biblioteca, func(l Libro) string {
		return fmt.Sprintf("%s por %s - %.1f/5.0", l.Titulo, l.Autor, l.Calificacion)
	})
	fmt.Println("Resúmenes:")
	for _, r := range resumenLibros {
		fmt.Printf("  %s\n", r)
	}

	
	fmt.Println("\n5. ORDENAMIENTO PERSONALIZABLE:")
	fmt.Println(strings.Repeat("-", 40))

	porCalificacionDesc := OrdenarLibros(biblioteca, func(a, b Libro) bool {
		return a.Calificacion > b.Calificacion
	})
	fmt.Println("Top 3 por calificación:")
	for i, l := range porCalificacionDesc[:3] {
		fmt.Printf("  %d. %s (%.1f)\n", i+1, l.Titulo, l.Calificacion)
	}

	porPaginasAsc := OrdenarLibros(biblioteca, func(a, b Libro) bool {
		return a.Paginas < b.Paginas
	})
	fmt.Println("Libros más cortos primero:")
	for i, l := range porPaginasAsc[:3] {
		fmt.Printf("  %d. %s (%d págs)\n", i+1, l.Titulo, l.Paginas)
	}

	
	fmt.Println("\n6. REDUCCIÓN:")
	fmt.Println(strings.Repeat("-", 40))

	totalPaginas := ReducirLibros(biblioteca, 0, func(acc float64, l Libro) float64 {
		return acc + float64(l.Paginas)
	})
	fmt.Printf("Total páginas en biblioteca: %.0f\n", totalPaginas)

	sumaCalificaciones := ReducirLibros(biblioteca, 0, func(acc float64, l Libro) float64 {
		return acc + l.Calificacion
	})
	promedio := sumaCalificaciones / float64(len(biblioteca))
	fmt.Printf("Calificación promedio: %.2f\n", promedio)

	
	fmt.Println("\n7. CLOSURES - CONTADORES:")
	fmt.Println(strings.Repeat("-", 40))

	contador1 := createCounter()
	contador2 := createCounter()
	fmt.Printf("Contador1: %d, %d, %d\n", contador1(), contador1(), contador1())
	fmt.Printf("Contador2: %d, %d\n", contador2(), contador2())
	fmt.Printf("Contador1 sigue: %d\n", contador1())

	incrementar, resetear := createCounterConReset()
	fmt.Printf("Con reset: %d, %d, %d\n", incrementar(), incrementar(), incrementar())
	resetear()
	fmt.Printf("Después del reset: %d\n", incrementar())

	
	fmt.Println("\n8. CLOSURES - ACUMULADOR DE CALIFICACIONES:")
	fmt.Println(strings.Repeat("-", 40))

	acumulador := CrearAcumuladorCalificaciones()
	for _, libro := range biblioteca[:4] {
		promActual := acumulador(libro.Calificacion)
		fmt.Printf("Agregado '%s' (%.1f) -> Promedio actual: %.2f\n",
			libro.Titulo, libro.Calificacion, promActual)
	}

	
	fmt.Println("\n9. PIPELINE COMPLETO:")
	fmt.Println(strings.Repeat("-", 40))

	
	filtroFinal := CombinarFiltros(
		CrearFiltroCalificacion(4.6),
		func(l Libro) bool { return l.Paginas < 600 },
	)
	resultado := FiltrarLibros(biblioteca, filtroFinal)
	resultado = OrdenarLibros(resultado, func(a, b Libro) bool {
		return a.Calificacion > b.Calificacion
	})
	resumenes := MapearLibros(resultado, func(l Libro) string {
		return fmt.Sprintf("★ %s [%s, %d págs] %.1f/5", l.Titulo, l.Genero, l.Paginas, l.Calificacion)
	})
	fmt.Println("Mejores libros (calif>=4.6, <600 págs), ordenados:")
	for _, r := range resumenes {
		fmt.Printf("  %s\n", r)
	}
}
