package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/JabberquackerWasTaken/SisDis/chat"
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
func enviarOrden(Lista []Orden) Orden {
	envio := Lista[0]
	Lista = append(Lista[:0], Lista[1:]...)
	/*Aca falta enviar envio a camiones por grpc*/
	//fmt.Println(envio)
	return envio
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
		me := chat.Orden{
			ID:        line[0],
			Producto:  line[1],
			Valor:     line[2],
			Tienda:    line[3],
			Destino:   line[4],
			Prioridad: "1",
		}

		response, _ := c.MandarOrden(context.Background(), &me)
		fmt.Println(response)
	}

	retail = append(retail[:0], retail[1:]...)
	ord := enviarOrden(retail)
	if ord.ID == "Quiero que deje te tirar que no lo estoy usando" {
		println("linea imposible")
	}
	ord = enviarOrden(pymes)

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
				fmt.Println("1.-Retail")
				ord = enviarOrden(retail)
				//response, err := c.MandarOrden(context.Background(), ord)
			} else {
				fmt.Println("No quedan ordenes de Retail en la lista.")
			}
		} else if strings.Compare(text, "2") == 0 {
			if len(retail) != 0 {
				fmt.Println("2.- Pymes")
				ord = enviarOrden(pymes)
				//response, err := c.MandarOrden(context.Background(), ord)
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
