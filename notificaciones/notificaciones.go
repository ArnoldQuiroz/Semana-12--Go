package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Notificador interface {
	EnviarNotificacion(destinatario, mensaje string) error
}

type ValidadorMensaje interface {
	ValidarMensaje(mensaje string) error
	ValidarDestinatario(destinatario string) error
}

type Rastreador interface {
	ObtenerEstado(id string) (string, error)
	ObtenerEstadisticas() map[string]int
}

type Logger interface {
	Log(nivel, mensaje string)
	LogError(error)
	LogInfo(string)
}

type NotificadorCompleto interface {
	Notificador
	ValidadorMensaje
}

type NotificadorAvanzado interface {
	Notificador
	ValidadorMensaje
	Rastreador
	Logger
}

type TipoNotificacion string

const (
	Email TipoNotificacion = "email"
	SMS   TipoNotificacion = "sms"
	Slack TipoNotificacion = "slack"
)

type EstadoNotificacion string

const (
	Pendiente EstadoNotificacion = "pendiente"
	Enviada   EstadoNotificacion = "enviada"
	Fallida   EstadoNotificacion = "fallida"
)

type RegistroNotificacion struct {
	ID           string
	Tipo         TipoNotificacion
	Destinatario string
	Mensaje      string
	Estado       EstadoNotificacion
	Timestamp    time.Time
	Intentos     int
	Error        string
}

type EmailNotificador struct {
	servidor  string
	puerto    int
	usuario   string
	password  string
	registros map[string]*RegistroNotificacion
}

func NuevoEmailNotificador(servidor string, puerto int, usuario, password string) *EmailNotificador {
	return &EmailNotificador{
		servidor:  servidor,
		puerto:    puerto,
		usuario:   usuario,
		password:  password,
		registros: make(map[string]*RegistroNotificacion),
	}
}

func (e *EmailNotificador) EnviarNotificacion(destinatario, mensaje string) error {
	if err := e.ValidarDestinatario(destinatario); err != nil {
		return err
	}
	if err := e.ValidarMensaje(mensaje); err != nil {
		return err
	}
	id := fmt.Sprintf("email_%d", time.Now().UnixNano())
	registro := &RegistroNotificacion{
		ID:           id,
		Tipo:         Email,
		Destinatario: destinatario,
		Mensaje:      mensaje,
		Estado:       Pendiente,
		Timestamp:    time.Now(),
		Intentos:     1,
	}
	e.registros[id] = registro
	e.LogInfo(fmt.Sprintf("Enviando email a %s", destinatario))
	time.Sleep(100 * time.Millisecond)
	if time.Now().UnixNano()%10 == 0 {
		registro.Estado = Fallida
		registro.Error = "Servidor SMTP no disponible"
		e.LogError(errors.New(registro.Error))
		return errors.New("fallo al enviar email")
	}
	registro.Estado = Enviada
	e.LogInfo(fmt.Sprintf("Email enviado exitosamente: %s", id))
	return nil
}

func (e *EmailNotificador) ValidarMensaje(mensaje string) error {
	if len(mensaje) == 0 {
		return errors.New("mensaje no puede estar vacío")
	}
	if len(mensaje) > 1000 {
		return errors.New("mensaje muy largo (máximo 1000 caracteres)")
	}
	return nil
}

func (e *EmailNotificador) ValidarDestinatario(destinatario string) error {
	if !strings.Contains(destinatario, "@") {
		return errors.New("email inválido: debe contener @")
	}
	if !strings.Contains(destinatario, ".") {
		return errors.New("email inválido: debe contener dominio")
	}
	return nil
}

func (e *EmailNotificador) ObtenerEstado(id string) (string, error) {
	if registro, existe := e.registros[id]; existe {
		return string(registro.Estado), nil
	}
	return "", errors.New("notificación no encontrada")
}

func (e *EmailNotificador) ObtenerEstadisticas() map[string]int {
	stats := map[string]int{"total": 0, "enviadas": 0, "fallidas": 0, "pendientes": 0}
	for _, registro := range e.registros {
		stats["total"]++
		switch registro.Estado {
		case Enviada:
			stats["enviadas"]++
		case Fallida:
			stats["fallidas"]++
		case Pendiente:
			stats["pendientes"]++
		}
	}
	return stats
}

func (e *EmailNotificador) Log(nivel, mensaje string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] EMAIL [%s]: %s\n", timestamp, nivel, mensaje)
}

func (e *EmailNotificador) LogError(err error) { e.Log("ERROR", err.Error()) }
func (e *EmailNotificador) LogInfo(mensaje string) { e.Log("INFO", mensaje) }

type SMSNotificador struct {
	apiKey    string
	proveedor string
	registros map[string]*RegistroNotificacion
}

func NuevoSMSNotificador(apiKey, proveedor string) *SMSNotificador {
	return &SMSNotificador{
		apiKey:    apiKey,
		proveedor: proveedor,
		registros: make(map[string]*RegistroNotificacion),
	}
}

func (s *SMSNotificador) EnviarNotificacion(destinatario, mensaje string) error {
	if err := s.ValidarDestinatario(destinatario); err != nil {
		return err
	}
	if err := s.ValidarMensaje(mensaje); err != nil {
		return err
	}
	id := fmt.Sprintf("sms_%d", time.Now().UnixNano())
	registro := &RegistroNotificacion{
		ID: id, Tipo: SMS, Destinatario: destinatario,
		Mensaje: mensaje, Estado: Pendiente, Timestamp: time.Now(), Intentos: 1,
	}
	s.registros[id] = registro
	s.LogInfo(fmt.Sprintf("Enviando SMS a %s via %s", destinatario, s.proveedor))
	time.Sleep(50 * time.Millisecond)
	if time.Now().UnixNano()%20 == 0 {
		registro.Estado = Fallida
		registro.Error = "Número no válido"
		s.LogError(errors.New(registro.Error))
		return errors.New("fallo al enviar SMS")
	}
	registro.Estado = Enviada
	s.LogInfo(fmt.Sprintf("SMS enviado exitosamente: %s", id))
	return nil
}

func (s *SMSNotificador) ValidarMensaje(mensaje string) error {
	if len(mensaje) == 0 {
		return errors.New("mensaje SMS no puede estar vacío")
	}
	if len(mensaje) > 160 {
		return errors.New("mensaje SMS muy largo (máximo 160 caracteres)")
	}
	return nil
}

func (s *SMSNotificador) ValidarDestinatario(destinatario string) error {
	if len(destinatario) < 10 {
		return errors.New("número de teléfono muy corto")
	}
	if !strings.HasPrefix(destinatario, "+") && !strings.HasPrefix(destinatario, "0") {
		return errors.New("número debe empezar con + o 0")
	}
	return nil
}

func (s *SMSNotificador) ObtenerEstado(id string) (string, error) {
	if registro, existe := s.registros[id]; existe {
		return string(registro.Estado), nil
	}
	return "", errors.New("SMS no encontrado")
}

func (s *SMSNotificador) ObtenerEstadisticas() map[string]int {
	stats := map[string]int{"total": 0, "enviados": 0, "fallidos": 0, "pendientes": 0}
	for _, r := range s.registros {
		stats["total"]++
		switch r.Estado {
		case Enviada:
			stats["enviados"]++
		case Fallida:
			stats["fallidos"]++
		case Pendiente:
			stats["pendientes"]++
		}
	}
	return stats
}

func (s *SMSNotificador) Log(nivel, mensaje string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] SMS [%s]: %s\n", timestamp, nivel, mensaje)
}

func (s *SMSNotificador) LogError(err error) { s.Log("ERROR", err.Error()) }
func (s *SMSNotificador) LogInfo(mensaje string) { s.Log("INFO", mensaje) }

type SlackNotificador struct {
	webhook string
	canal   string
}

func NuevoSlackNotificador(webhook, canal string) *SlackNotificador {
	return &SlackNotificador{webhook: webhook, canal: canal}
}

func (sl *SlackNotificador) EnviarNotificacion(destinatario, mensaje string) error {
	fmt.Printf("Slack -> Canal: %s | Usuario: %s | Mensaje: %s\n", sl.canal, destinatario, mensaje)
	time.Sleep(10 * time.Millisecond)
	return nil
}

type ServicioNotificaciones struct {
	notificadores []Notificador
	logger        Logger
}

func NuevoServicioNotificaciones() *ServicioNotificaciones {
	return &ServicioNotificaciones{
		notificadores: make([]Notificador, 0),
	}
}

func (sn *ServicioNotificaciones) EstablecerLogger(logger Logger) {
	sn.logger = logger
}

func (sn *ServicioNotificaciones) AgregarNotificador(notificador Notificador) {
	sn.notificadores = append(sn.notificadores, notificador)
	if sn.logger != nil {
		sn.logger.LogInfo(fmt.Sprintf("Notificador agregado: %T", notificador))
	}
}

func (sn *ServicioNotificaciones) EnviarATodos(destinatario, mensaje string) map[string]error {
	resultados := make(map[string]error)
	if sn.logger != nil {
		sn.logger.LogInfo(fmt.Sprintf("Enviando a %d notificadores", len(sn.notificadores)))
	}
	for _, notificador := range sn.notificadores {
		tipo := fmt.Sprintf("%T", notificador)
		err := notificador.EnviarNotificacion(destinatario, mensaje)
		resultados[tipo] = err
		if sn.logger != nil {
			if err != nil {
				sn.logger.LogError(fmt.Errorf("%s falló: %v", tipo, err))
			} else {
				sn.logger.LogInfo(fmt.Sprintf("%s éxito", tipo))
			}
		}
	}
	return resultados
}

func (sn *ServicioNotificaciones) EnviarConValidacion(destinatario, mensaje string) map[string]error {
	resultados := make(map[string]error)
	for _, notificador := range sn.notificadores {
		tipo := fmt.Sprintf("%T", notificador)
		if validador, implementa := notificador.(ValidadorMensaje); implementa {
			if err := validador.ValidarMensaje(mensaje); err != nil {
				resultados[tipo] = fmt.Errorf("validación falló: %v", err)
				continue
			}
			if err := validador.ValidarDestinatario(destinatario); err != nil {
				resultados[tipo] = fmt.Errorf("destinatario inválido: %v", err)
				continue
			}
		}
		err := notificador.EnviarNotificacion(destinatario, mensaje)
		resultados[tipo] = err
	}
	return resultados
}

func ProbarNotificador(n Notificador, destinatario, mensaje string) {
	fmt.Printf("\nProbando %T:\n", n)
	fmt.Println("  Enviando:", mensaje)
	err := n.EnviarNotificacion(destinatario, mensaje)
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Println("  Enviado correctamente")
	}
}

func AnalizarCapacidadesNotificador(n Notificador) {
	fmt.Printf("\nAnalizando capacidades de %T:\n", n)
	capacidades := []string{}
	if _, ok := n.(Notificador); ok {
		capacidades = append(capacidades, "Notificador (envío básico)")
	}
	if _, ok := n.(ValidadorMensaje); ok {
		capacidades = append(capacidades, "ValidadorMensaje (validación)")
	}
	if _, ok := n.(Rastreador); ok {
		capacidades = append(capacidades, "Rastreador (seguimiento)")
	}
	if _, ok := n.(Logger); ok {
		capacidades = append(capacidades, "Logger (registro de eventos)")
	}
	if _, ok := n.(NotificadorCompleto); ok {
		capacidades = append(capacidades, "NotificadorCompleto")
	}
	if _, ok := n.(NotificadorAvanzado); ok {
		capacidades = append(capacidades, "NotificadorAvanzado")
	}
	for _, c := range capacidades {
		fmt.Printf("  - %s\n", c)
	}
}

func main() {
	fmt.Println("SISTEMA DE NOTIFICACIONES - INTERFACES EN ACCIÓN")
	fmt.Println(strings.Repeat("=", 60))

	email := NuevoEmailNotificador("smtp.gmail.com", 587, "app@empresa.com", "password")
	sms := NuevoSMSNotificador("api-key-123", "Twilio")
	slack := NuevoSlackNotificador("https://hooks.slack.com/...", "#general")

	servicio := NuevoServicioNotificaciones()
	servicio.EstablecerLogger(email)
	servicio.AgregarNotificador(email)
	servicio.AgregarNotificador(sms)
	servicio.AgregarNotificador(slack)

	
	fmt.Println("\n1. POLIMORFISMO BÁSICO:")
	fmt.Println(strings.Repeat("-", 40))
	notificadores := []Notificador{email, sms, slack}
	for _, n := range notificadores {
		ProbarNotificador(n, "usuario@ejemplo.com", "¡Hola desde Go!")
	}

	
	fmt.Println("\n2. TYPE ASSERTIONS Y CAPACIDADES:")
	fmt.Println(strings.Repeat("-", 40))
	for _, n := range notificadores {
		AnalizarCapacidadesNotificador(n)
	}

	
	fmt.Println("\n3. TYPE SWITCH EN ACCIÓN:")
	fmt.Println(strings.Repeat("-", 40))
	for _, n := range notificadores {
		switch notificador := n.(type) {
		case *EmailNotificador:
			fmt.Printf("EmailNotificador - Servidor: %s:%d\n", notificador.servidor, notificador.puerto)
			notificador.EnviarNotificacion("test@test.com", "mensaje email")
		case *SMSNotificador:
			fmt.Printf("SMSNotificador - Proveedor: %s\n", notificador.proveedor)
			notificador.EnviarNotificacion("+54911234567", "mensaje sms")
		case *SlackNotificador:
			fmt.Printf("SlackNotificador - Canal: %s\n", notificador.canal)
			notificador.EnviarNotificacion("usuario", "mensaje slack")
		}
	}

	
	fmt.Println("\n4. INTERFACES COMPUESTAS:")
	fmt.Println(strings.Repeat("-", 40))
	verificar := func(n Notificador) {
		nombre := fmt.Sprintf("%T", n)
		if completo, ok := n.(NotificadorCompleto); ok {
			fmt.Printf("%s implementa NotificadorCompleto\n", nombre)
			completo.ValidarMensaje("test")
		} else {
			fmt.Printf("%s NO implementa NotificadorCompleto\n", nombre)
		}
		if _, ok := n.(NotificadorAvanzado); ok {
			fmt.Printf("%s implementa NotificadorAvanzado\n", nombre)
		} else {
			fmt.Printf("%s NO implementa NotificadorAvanzado\n", nombre)
		}
	}
	for _, n := range notificadores {
		verificar(n)
	}

	
	fmt.Println("\n5. SERVICIO CON MÚLTIPLES NOTIFICADORES:")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("Enviando a TODOS los notificadores:")
	resultados := servicio.EnviarATodos("admin@empresa.com", "Sistema iniciado correctamente")
	for tipo, err := range resultados {
		if err != nil {
			fmt.Printf("  %s: Error - %v\n", tipo, err)
		} else {
			fmt.Printf("  %s: Éxito\n", tipo)
		}
	}

	fmt.Println("\nEjemplo completado!")
}
