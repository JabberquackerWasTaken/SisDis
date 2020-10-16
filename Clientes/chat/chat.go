package chat

import (
	"log"

	"golang.org/x/net/context"
)

// Orden struct recibe las ordenes
type Orden struct {
	ID        string `json:"id"`
	Producto  string `json:"producto"`
	Valor     string `json:"valor"`
	Tienda    string `json:"tienda"`
	Destino   string `json:"destino"`
	Prioridad string `json:"prioridad"`
}

// Server es un struct de server
type Server struct {
}

//MandarOrden es la funcion que usamos para mandar mensajes a los camiones
func (s *Server) MandarOrden(ctx context.Context, message *Orden) (*Message, error) {
	log.Printf("Orden recibida")
	return &Message{Body: "Orden enviada"}, nil
}

//SayHello envia mensajes entre servidor-cliente
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
