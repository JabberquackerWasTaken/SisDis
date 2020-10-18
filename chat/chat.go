package chat

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/net/context"
)

// Orden struct recibe las ordenes

// Server es un struct de server
type Server struct {
	Ordenes  []string
	Finanzas []string
}

//MandarOrden es la funcion que usamos para mandar mensajes a los camiones
func (s *Server) MandarOrden(ctx context.Context, message *Orden) (*Message, error) {
	fmt.Println("Orden recibida")
	return &Message{Body: "Orden enviada"}, nil
}

//SayHola envia mensajes entre servidor-cliente
func (s *Server) SayHola(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	respuesta := Message{
		Body: "iniciando",
	}
	if strings.Compare(message.Body, "Hay entregas?") == 0 {
		if len(s.Ordenes) == 0 {
			//no
			respuesta = Message{
				Body: "Nada@Nada@Nada@Nada@Nada@Nada",
			}
		} else {
			//si
			envio := s.Ordenes[0]

			respuesta = Message{
				Body: envio,
			}
			s.Ordenes = append(s.Ordenes[:0], s.Ordenes[1:]...)
		}
	} else if strings.Compare(message.Body, "Finanzas") == 0 {

		envio := s.Finanzas[0]

		respuesta = Message{
			Body: envio,
		}
		s.Finanzas = append(s.Finanzas[:0], s.Finanzas[1:]...)

	} else if strings.Compare(message.Body, "Largo") == 0 {
		respuesta = Message{
			Body: strconv.Itoa(len(s.Finanzas)),
		}
	} else {
		romper := strings.SplitN(message.Body, "@", 6)
		if strings.Compare(romper[3], "1") == 0 {
			s.Finanzas = append(s.Finanzas, message.Body)
			respuesta = Message{
				Body: "recibido",
			}
		} else if strings.Compare(romper[3], "0") == 0 {
			s.Finanzas = append(s.Finanzas, message.Body)
			respuesta = Message{
				Body: "recibido",
			}
		} else {
			s.Ordenes = append(s.Ordenes, message.Body)
			respuesta = Message{
				Body: "Orden recibida",
			}
		}
	}
	return &respuesta, nil
}
