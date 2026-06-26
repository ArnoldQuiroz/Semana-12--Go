package main

import (
	"fmt"
	"math"
)

type Figura interface {
	Area() float64
	Perimetro() float64
	Describir() string
}

type FiguraColoreada interface {
	Figura
	Color() string
}

type Rectangulo struct {
	Ancho  float64
	Alto   float64
	ColorRGB string
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func (r Rectangulo) Perimetro() float64 {
	return 2 * (r.Ancho + r.Alto)
}

func (r Rectangulo) Describir() string {
	return fmt.Sprintf("Rectángulo %.1fx%.1f", r.Ancho, r.Alto)
}

func (r Rectangulo) Color() string {
	return r.ColorRGB
}

type Circulo struct {
	Radio    float64
	ColorRGB string
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

func (c Circulo) Describir() string {
	return fmt.Sprintf("Círculo radio=%.1f", c.Radio)
}

func (c Circulo) Color() string {
	return c.ColorRGB
}

type Triangulo struct {
	Base   float64
	Altura float64
	LadoA  float64
	LadoB  float64
	LadoC  float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

func (t Triangulo) Perimetro() float64 {
	return t.LadoA + t.LadoB + t.LadoC
}

func (t Triangulo) Describir() string {
	return fmt.Sprintf("Triángulo base=%.1f altura=%.1f", t.Base, t.Altura)
}

func imprimirInfo(f Figura) {
	fmt.Printf("%s → Área: %.2f | Perímetro: %.2f\n",
		f.Describir(), f.Area(), f.Perimetro())
}

func areaTotal(figuras []Figura) float64 {
	total := 0.0
	for _, f := range figuras {
		total += f.Area()
	}
	return total
}

func main() {
	r := Rectangulo{Ancho: 5, Alto: 3, ColorRGB: "rojo"}
	c := Circulo{Radio: 4, ColorRGB: "azul"}
	t := Triangulo{Base: 6, Altura: 4, LadoA: 5, LadoB: 5, LadoC: 6}

	
	figuras := []Figura{r, c, t}
	fmt.Println("=== Información de figuras ===")
	for _, f := range figuras {
		imprimirInfo(f)
	}

	fmt.Printf("\nÁrea total: %.2f\n", areaTotal(figuras))

	
	fmt.Println("\n=== Type Assertions ===")
	for _, f := range figuras {
		if fc, ok := f.(FiguraColoreada); ok {
			fmt.Printf("%s tiene color: %s\n", fc.Describir(), fc.Color())
		} else {
			nombre := fmt.Sprintf("%T", f)
			fmt.Printf("%s no implementa FiguraColoreada\n", nombre)
		}
	}

	
	fmt.Println("\n=== Type Switch ===")
	for _, f := range figuras {
		switch fig := f.(type) {
		case Rectangulo:
			fmt.Printf("Es un Rectángulo: %dx%d\n", int(fig.Ancho), int(fig.Alto))
		case Circulo:
			fmt.Printf("Es un Círculo con radio %.0f\n", fig.Radio)
		case Triangulo:
			fmt.Printf("Es un Triángulo con base %.0f\n", fig.Base)
		default:
			fmt.Println("Tipo desconocido")
		}
	}

	
	var f Figura
	f = r
	fmt.Printf("\nTipo de f: %T\n", f)
	f = c
	fmt.Printf("Tipo de f: %T\n", f)
}
