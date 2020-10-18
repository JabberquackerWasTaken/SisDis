package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/JabberquackerWasTaken/SisDis/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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

//enviarOrden es una funcion que remueve el primer item de la cola y lo envia a los camiones.
func enviarOrden(Lista []Orden) chat.Message {
	envio := Lista[0]
	Lista = append(Lista[:0], Lista[1:]...)
	respuesta := chat.Message{
		Body: envio.ID + "@" + envio.Producto + "@" + envio.Valor + "@" + envio.Tienda + "@" + envio.Destino + "@" + envio.Prioridad,
	}
	return respuesta
}

func main() {
	var conec *grpc.ClientConn
	conec, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conec.Close()

	c := chat.NewChatServiceClient(conec)
	//Pymes
	csvFile, _ := os.Open("pymes.csv")
	csvFile2, _ := os.Open("retail.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var pymes []Orden
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		pymes = append(pymes, Orden{
			ID:        line[0],
			Producto:  line[1],
			Valor:     line[2],
			Tienda:    line[3],
			Destino:   line[4],
			Prioridad: line[5],
		})
	}
	pymes = append(pymes[:0], pymes[1:]...)
	reader2 := csv.NewReader(bufio.NewReader(csvFile2))
	var retail []Orden
	for {
		line, error := reader2.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		retail = append(retail, Orden{
			ID:        line[0],
			Producto:  line[1],
			Valor:     line[2],
			Tienda:    line[3],
			Destino:   line[4],
			Prioridad: "1",
		})
	}
	retail = append(retail[:0], retail[1:]...)
	var ordenEnvio = chat.Message{
		Body: "iniciando esto",
	}
	var request = bufio.NewReader(os.Stdin)
	for {
		fmt.Println("----------------------------")
		fmt.Println("Ingrese una opcion de Orden:")
		fmt.Println("1.-Retail")
		fmt.Println("2.-Pymes")
		fmt.Println("3.-Salir")
		fmt.Println("----------------------------")
		fmt.Print("Opcion: ")
		text, _ := request.ReadString('\n')
		text = strings.ToLower(strings.Trim(text, " \r\n"))
		if strings.Compare(text, "1") == 0 {
			if len(retail) != 0 {
				fmt.Println("Opcion de Retail seleccionada")
				ordenEnvio = enviarOrden(retail)
				response, err := c.SayHello(context.Background(), &ordenEnvio)
				if err != nil {
					log.Fatalf("Error when calling SayHello: %s", err)
				}
				fmt.Println(response)
			} else {
				fmt.Println("No quedan ordenes de Retail en la lista.")
			}
		} else if strings.Compare(text, "2") == 0 {
			if len(retail) != 0 {
				fmt.Println("Opcion de Pymes seleccionada")
				ordenEnvio = enviarOrden(pymes)
				response, err := c.SayHello(context.Background(), &ordenEnvio)
				if err != nil {
					log.Fatalf("Error when calling SayHello: %s", err)
				}
				fmt.Println(response)
			} else {
				fmt.Println("No quedan ordenes de Pymes en la lista.")
			}
		} else if strings.Compare(text, "3") == 0 {
			break
		} else {
			fmt.Println("----------------------------")
			fmt.Println("Opcion invalida")
		}
	}
	fmt.Println("Gracias por preferirnos")
}
