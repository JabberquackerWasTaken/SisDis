package chat

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
)

// Orden struct recibe las ordenes

// Server es un struct de server
type Server struct {
}

//MandarOrden es la funcion que usamos para mandar mensajes a los camiones
func (s *Server) MandarOrden(ctx context.Context, message *Orden) (*Message, error) {
	fmt.Println("Orden recibida")
	return &Message{Body: "Orden enviada"}, nil
}

//SayHello envia mensajes entre servidor-cliente
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
