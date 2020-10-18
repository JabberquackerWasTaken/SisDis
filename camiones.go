package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/JabberquackerWasTaken/SisDis/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Orden struct recibe las ordenes
type Orden struct {
	ID        string
	Producto  string
	Valor     string
	Tienda    string
	Destino   string
	Prioridad string
}

//Camion es el struct para los camiones
type Camion struct {
	Orden1   Orden
	Orden2   Orden
	intento  int
	intento2 int
}

func main() {
	//	var Aux Orden
	var conec *grpc.ClientConn
	conec, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conec.Close()

	c := chat.NewChatServiceClient(conec)
	var i int
	//	var Cam Camion

	fmt.Println("----------------------------")
	fmt.Println("Los Camiones revisaran que no hayan nuevas entregas cada 7 segundos")
	for {
		for i = 0; i < 7; i++ {
			time.Sleep(time.Second)
		}
		message := chat.Message{
			Body: "Hay entregas?",
		}
		response, err := c.SayHola(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling server: %s", err)
		}
		res1 := strings.SplitN(response.Body, "@", 6)
		/*Aux = append(Aux, Orden{
			ID:        res1[0],
			Producto:  res1[1],
			Valor:     res1[2],
			Tienda:    res1[3],
			Destino:   res1[4],
			Prioridad: res1[5],
		})*/
		fmt.Println(res1)
	}
}
