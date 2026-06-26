package main

import "fmt"

type Producto struct {
	ID     int
	Nombre string
	Precio float64
	Stock  int
}

func (p Producto) ObtenerInfo() string {
	return fmt.Sprintf("[%d] %s - S/%.2f (Stock: %d)", p.ID, p.Nombre, p.Precio, p.Stock)
}

func (p Producto) EstaDisponible() bool {
	return p.Stock > 0
}

func (p *Producto) AplicarDescuento(porcentaje float64) {
	p.Precio = p.Precio * (1 - porcentaje/100)
}

func (p *Producto) ActualizarStock(cantidad int) error {
	if p.Stock+cantidad < 0 {
		return fmt.Errorf("stock insuficiente: solo hay %d unidades", p.Stock)
	}
	p.Stock += cantidad
	return nil
}

func (p *Producto) ActualizarPrecio(nuevoPrecio float64) error {
	if nuevoPrecio <= 0 {
		return fmt.Errorf("el precio debe ser mayor a cero")
	}
	p.Precio = nuevoPrecio
	return nil
}

type Carrito struct {
	Items    []Producto
	Total    float64
}

func (c Carrito) ContarItems() int {
	return len(c.Items)
}

func (c *Carrito) AgregarProducto(p Producto) {
	c.Items = append(c.Items, p)
	c.Total += p.Precio
}

func (c *Carrito) Vaciar() {
	c.Items = []Producto{}
	c.Total = 0
}

func main() {
	p1 := Producto{ID: 1, Nombre: "Laptop", Precio: 2500.00, Stock: 5}

	
	fmt.Println(p1.ObtenerInfo())
	fmt.Println("Disponible:", p1.EstaDisponible())

	
	p1.AplicarDescuento(10)
	fmt.Println("Después del descuento:", p1.ObtenerInfo())

	err := p1.ActualizarStock(-3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Stock actualizado:", p1.Stock)
	}

	
	err = p1.ActualizarStock(-10)
	if err != nil {
		fmt.Println("Error esperado:", err)
	}

	
	p2 := Producto{ID: 2, Nombre: "Mouse", Precio: 50.00, Stock: 10}
	copia := p2
	copia.Precio = 999 
	fmt.Printf("Original: S/%.2f | Copia: S/%.2f\n", p2.Precio, copia.Precio)

	
	carrito := Carrito{}
	carrito.AgregarProducto(p1)
	carrito.AgregarProducto(p2)
	fmt.Printf("Carrito: %d items, Total: S/%.2f\n", carrito.ContarItems(), carrito.Total)

	carrito.Vaciar()
	fmt.Println("Carrito vaciado, items:", carrito.ContarItems())
}
