package chat

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/net/context"
)

// Orden struct recibe las ordenes

// Server es un struct de server con 2 Listas de strings
type Server struct {
	Ordenes  []string
	Finanzas []string
}

//MandarOrden es la funcion que finalmente no usamos en este trabajo por problemas de tiempo/el chat.pb.go decidio no reconocerla.
func (s *Server) MandarOrden(ctx context.Context, message *Orden) (*Message, error) {
	fmt.Println("Orden recibida")
	return &Message{Body: "Orden enviada"}, nil
}

//SayHola envia mensajes entre servidor-cliente y siempre es un chat.Message, Lo que el string dire depende de lo que el cliente le pide
func (s *Server) SayHola(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	respuesta := Message{
		Body: "iniciando",
	}
	//Camiones preguntando si hay nuevas entregas o no
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
		//Finanzas pidiendo un reporte
		envio := s.Finanzas[0]

		respuesta = Message{
			Body: envio,
		}
		s.Finanzas = append(s.Finanzas[:0], s.Finanzas[1:]...)

	} else if strings.Compare(message.Body, "Largo") == 0 {
		//Finanzas preguntando cuantos reportes hay
		respuesta = Message{
			Body: strconv.Itoa(len(s.Finanzas)),
		}
	} else {
		/*Rompemos el string y lo guardamos en la Lista ordenes o Reportes,
		la manera que esto funciona es:
		En el 3er valor las ordenes solo pueden tener el nombre de la "tienda-Letra" o "pyme"
		mientras que los reportes en el tercer valor solo pueden tener "1" o "0"
		*/
		romper := strings.SplitN(message.Body, "@", 6)
		if strings.Compare(romper[3], "1") == 0 {
			s.Finanzas = append(s.Finanzas, message.Body)
			respuesta = Message{
				Body: "Reporte Enviado",
			}
		} else if strings.Compare(romper[3], "0") == 0 {
			s.Finanzas = append(s.Finanzas, message.Body)
			respuesta = Message{
				Body: "Reporte Enviado",
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
