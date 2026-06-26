package main

import (
	"fmt"
	"strings"
	"time"
)

type Libro struct {
	ID       int
	Titulo   string
	Autor    string
	ISBN     string
	Paginas  int
	Prestado bool
}

type Usuario struct {
	ID       int
	Nombre   string
	Email    string
	Telefono string
	Activo   bool
}

type Prestamo struct {
	ID              int
	LibroID         int
	UsuarioID       int
	FechaPrestamo   time.Time
	FechaDevolucion time.Time
	Devuelto        bool
}

func (l Libro) ObtenerInfo() string {
	estado := "Disponible"
	if l.Prestado {
		estado = "Prestado"
	}
	return fmt.Sprintf("[%d] %s por %s - %s", l.ID, l.Titulo, l.Autor, estado)
}

func (l Libro) EsPrestable() bool {
	return !l.Prestado && l.Paginas > 0
}

func (l Libro) EsLibroGrande() bool {
	return l.Paginas > 300
}

func (u Usuario) ObtenerResumen() string {
	estado := "Inactivo"
	if u.Activo {
		estado = "Activo"
	}
	return fmt.Sprintf("%s (%s) - %s", u.Nombre, u.Email, estado)
}

func (u Usuario) PuedePrestar() bool {
	return u.Activo && u.Email != "" && u.Nombre != ""
}

func (l *Libro) Prestar() error {
	if l.Prestado {
		return fmt.Errorf("el libro '%s' ya está prestado", l.Titulo)
	}
	if l.Paginas <= 0 {
		return fmt.Errorf("el libro '%s' no es válido", l.Titulo)
	}
	l.Prestado = true
	return nil
}

func (l *Libro) Devolver() error {
	if !l.Prestado {
		return fmt.Errorf("el libro '%s' no está prestado", l.Titulo)
	}
	l.Prestado = false
	return nil
}

func (l *Libro) ActualizarInfo(titulo, autor string, paginas int) error {
	if titulo == "" || autor == "" {
		return fmt.Errorf("título y autor no pueden estar vacíos")
	}
	if paginas <= 0 {
		return fmt.Errorf("número de páginas debe ser positivo")
	}
	l.Titulo = titulo
	l.Autor = autor
	l.Paginas = paginas
	return nil
}

func (u *Usuario) Activar() {
	u.Activo = true
}

func (u *Usuario) Desactivar() {
	u.Activo = false
}

func (u *Usuario) ActualizarContacto(email, telefono string) error {
	if !strings.Contains(email, "@") {
		return fmt.Errorf("email inválido: %s", email)
	}
	u.Email = email
	u.Telefono = telefono
	return nil
}

type Biblioteca struct {
	Nombre    string
	Direccion string
	Libros    []Libro
	Usuarios  []Usuario
	Prestamos []Prestamo
	proximoID int
}

func NuevaBiblioteca(nombre, direccion string) *Biblioteca {
	return &Biblioteca{
		Nombre:    nombre,
		Direccion: direccion,
		Libros:    make([]Libro, 0),
		Usuarios:  make([]Usuario, 0),
		Prestamos: make([]Prestamo, 0),
		proximoID: 1,
	}
}

func (b *Biblioteca) AgregarLibro(titulo, autor, isbn string, paginas int) (*Libro, error) {
	if titulo == "" || autor == "" {
		return nil, fmt.Errorf("título y autor son obligatorios")
	}
	for _, libro := range b.Libros {
		if libro.ISBN == isbn && isbn != "" {
			return nil, fmt.Errorf("ya existe un libro con ISBN: %s", isbn)
		}
	}
	libro := Libro{
		ID:       b.proximoID,
		Titulo:   titulo,
		Autor:    autor,
		ISBN:     isbn,
		Paginas:  paginas,
		Prestado: false,
	}
	b.Libros = append(b.Libros, libro)
	b.proximoID++
	return &libro, nil
}

func (b *Biblioteca) RegistrarUsuario(nombre, email, telefono string) (*Usuario, error) {
	if nombre == "" || email == "" {
		return nil, fmt.Errorf("nombre y email son obligatorios")
	}
	if !strings.Contains(email, "@") {
		return nil, fmt.Errorf("email inválido: %s", email)
	}
	for _, u := range b.Usuarios {
		if u.Email == email {
			return nil, fmt.Errorf("ya existe un usuario con email: %s", email)
		}
	}
	usuario := Usuario{
		ID:       b.proximoID,
		Nombre:   nombre,
		Email:    email,
		Telefono: telefono,
		Activo:   true,
	}
	b.Usuarios = append(b.Usuarios, usuario)
	b.proximoID++
	return &usuario, nil
}

func (b Biblioteca) BuscarLibro(id int) *Libro {
	for i, libro := range b.Libros {
		if libro.ID == id {
			return &b.Libros[i]
		}
	}
	return nil
}

func (b Biblioteca) BuscarUsuario(id int) *Usuario {
	for i, usuario := range b.Usuarios {
		if usuario.ID == id {
			return &b.Usuarios[i]
		}
	}
	return nil
}

func (b *Biblioteca) PrestarLibro(libroID, usuarioID int) error {
	libro := b.BuscarLibro(libroID)
	if libro == nil {
		return fmt.Errorf("libro con ID %d no encontrado", libroID)
	}
	usuario := b.BuscarUsuario(usuarioID)
	if usuario == nil {
		return fmt.Errorf("usuario con ID %d no encontrado", usuarioID)
	}
	if !usuario.PuedePrestar() {
		return fmt.Errorf("el usuario %s no puede realizar préstamos", usuario.Nombre)
	}
	if !libro.EsPrestable() {
		return fmt.Errorf("el libro '%s' no se puede prestar", libro.Titulo)
	}
	if err := libro.Prestar(); err != nil {
		return err
	}
	prestamo := Prestamo{
		ID:              b.proximoID,
		LibroID:         libroID,
		UsuarioID:       usuarioID,
		FechaPrestamo:   time.Now(),
		FechaDevolucion: time.Now().AddDate(0, 0, 14),
		Devuelto:        false,
	}
	b.Prestamos = append(b.Prestamos, prestamo)
	b.proximoID++
	return nil
}

func (b *Biblioteca) DevolverLibro(libroID int) error {
	libro := b.BuscarLibro(libroID)
	if libro == nil {
		return fmt.Errorf("libro con ID %d no encontrado", libroID)
	}
	var prestamoActivo *Prestamo
	for i := range b.Prestamos {
		if b.Prestamos[i].LibroID == libroID && !b.Prestamos[i].Devuelto {
			prestamoActivo = &b.Prestamos[i]
			break
		}
	}
	if prestamoActivo == nil {
		return fmt.Errorf("no se encontró préstamo activo para el libro '%s'", libro.Titulo)
	}
	if err := libro.Devolver(); err != nil {
		return err
	}
	prestamoActivo.Devuelto = true
	return nil
}

func (b Biblioteca) ObtenerEstadisticas() string {
	totalLibros := len(b.Libros)
	librosPrestados := 0
	usuariosActivos := 0
	prestamosActivos := 0

	for _, libro := range b.Libros {
		if libro.Prestado {
			librosPrestados++
		}
	}
	for _, usuario := range b.Usuarios {
		if usuario.Activo {
			usuariosActivos++
		}
	}
	for _, prestamo := range b.Prestamos {
		if !prestamo.Devuelto {
			prestamosActivos++
		}
	}
	return fmt.Sprintf(
		"Estadísticas de %s:\n  Total libros: %d\n  Libros prestados: %d\n  Libros disponibles: %d\n  Usuarios activos: %d\n  Préstamos activos: %d",
		b.Nombre, totalLibros, librosPrestados, totalLibros-librosPrestados, usuariosActivos, prestamosActivos,
	)
}

func (b Biblioteca) ListarLibrosDisponibles() {
	fmt.Println("\nLibros Disponibles:")
	fmt.Println(strings.Repeat("=", 50))
	disponibles := 0
	for _, libro := range b.Libros {
		if !libro.Prestado {
			fmt.Printf("  %s\n", libro.ObtenerInfo())
			if libro.EsLibroGrande() {
				fmt.Printf("    Libro extenso (%d páginas)\n", libro.Paginas)
			}
			disponibles++
		}
	}
	if disponibles == 0 {
		fmt.Println("  No hay libros disponibles")
	}
}

func main() {
	fmt.Println("SISTEMA DE BIBLIOTECA - DEMO PRÁCTICA")
	fmt.Println(strings.Repeat("=", 50))

	
	biblioteca := NuevaBiblioteca("Biblioteca Central", "Av. Principal 123")
	fmt.Printf("\nBiblioteca creada: %s\n", biblioteca.Nombre)

	
	fmt.Println("\nAgregando libros...")
	libros := []struct {
		titulo, autor, isbn string
		paginas             int
	}{
		{"El Quijote", "Miguel de Cervantes", "978-84-376-0494-7", 863},
		{"Cien Años de Soledad", "Gabriel García Márquez", "978-84-376-0495-4", 471},
		{"Go Programming", "Alan Donovan", "978-0-13-419044-0", 380},
		{"Clean Code", "Robert Martin", "978-0-13-235088-4", 464},
	}
	for _, l := range libros {
		libro, err := biblioteca.AgregarLibro(l.titulo, l.autor, l.isbn, l.paginas)
		if err != nil {
			fmt.Printf("  Error: %v\n", err)
		} else {
			fmt.Printf("  Agregado: %s\n", libro.ObtenerInfo())
		}
	}

	
	fmt.Println("\nRegistrando usuarios...")
	usuarios := []struct{ nombre, email, telefono string }{
		{"Ana García", "ana.garcia@email.com", "555-0101"},
		{"Carlos López", "carlos.lopez@email.com", "555-0102"},
		{"María Rodríguez", "maria.rodriguez@email.com", "555-0103"},
	}
	for _, u := range usuarios {
		usuario, err := biblioteca.RegistrarUsuario(u.nombre, u.email, u.telefono)
		if err != nil {
			fmt.Printf("  Error: %v\n", err)
		} else {
			fmt.Printf("  Registrado: %s\n", usuario.ObtenerResumen())
		}
	}

	
	fmt.Println("\nRealizando préstamos...")
	prestamos := []struct{ libroID, usuarioID int }{
		{1, 1}, 
		{3, 2}, 
		{2, 3}, 
	}
	for _, p := range prestamos {
		err := biblioteca.PrestarLibro(p.libroID, p.usuarioID)
		if err != nil {
			fmt.Printf("  Error en préstamo: %v\n", err)
		} else {
			libro := biblioteca.BuscarLibro(p.libroID)
			usuario := biblioteca.BuscarUsuario(p.usuarioID)
			fmt.Printf("  %s prestó '%s'\n", usuario.Nombre, libro.Titulo)
		}
	}

	
	biblioteca.ListarLibrosDisponibles()

	
	fmt.Println("\nDevolviendo libro...")
	err := biblioteca.DevolverLibro(1)
	if err != nil {
		fmt.Printf("  Error en devolución: %v\n", err)
	} else {
		fmt.Println("  El Quijote devuelto correctamente")
	}

	
	fmt.Println("\n" + biblioteca.ObtenerEstadisticas())

	
	fmt.Println("\nDEMO: Diferencia entre receptores")
	fmt.Println(strings.Repeat("=", 50))
	libro := biblioteca.BuscarLibro(4)
	fmt.Printf("Estado inicial: %s\n", libro.ObtenerInfo())
	err = libro.Prestar()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Después del préstamo: %s\n", libro.ObtenerInfo())
	}
	fmt.Printf("¿Es prestable?: %v\n", libro.EsPrestable())
	fmt.Printf("¿Es libro grande?: %v\n", libro.EsLibroGrande())
}
