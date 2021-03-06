package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/JabberquackerWasTaken/SisDis/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var Lista []string
	var conec *grpc.ClientConn
	var repeticiones int
	total := 0.0
	conec, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conec.Close()

	c := chat.NewChatServiceClient(conec)
	var request = bufio.NewReader(os.Stdin)
	for {
		fmt.Println("----------------------------")
		fmt.Println("Ingrese una opcion de Orden:")
		fmt.Println("1.-Numero de Ordenes")
		fmt.Println("2.-Obtener Ganancias")
		fmt.Println("3.-Salir")
		fmt.Println("----------------------------")
		fmt.Print("Opcion: ")
		text, _ := request.ReadString('\n')
		text = strings.ToLower(strings.Trim(text, " \r\n"))

		if strings.Compare(text, "1") == 0 {
			//honestamente esta opcion es bastante inutil y solo sirve por si la persona no quiere perder el tiempo calculando nuevas ordenes que o existen
			message := chat.Message{
				Body: "Largo",
			}
			response, err := c.SayHola(context.Background(), &message)
			if err != nil {
				log.Fatalf("Error when calling server: %s", err)
			}
			fmt.Println(response.Body)
		} else if strings.Compare(text, "2") == 0 {
			//Le pido al server el largo de la lista de Reportes
			message := chat.Message{
				Body: "Largo",
			}
			response, err := c.SayHola(context.Background(), &message)
			if err != nil {
				log.Fatalf("Error when calling server: %s", err)
			}
			repeticiones, _ = strconv.Atoi(response.Body)
			//Pido los reportes uno por uno y los guardo en la lista de reportes
			for i := 0; i < repeticiones; i++ {
				message := chat.Message{
					Body: "Finanzas",
				}
				response, err := c.SayHola(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling server: %s", err)
				}
				Lista = append(Lista, response.Body)
			}
			fmt.Println("Los nuevos Reportes obtenidos son:")
			//Separo los string y calculo cuanto gane por cada opcion
			for i := 0; i < repeticiones; i++ {
				res := strings.SplitN(Lista[0], "@", 8)

				fmt.Println(res)
				intentos, _ := strconv.Atoi(res[5])
				prioridad, _ := strconv.Atoi(res[3])
				valor, _ := strconv.Atoi(res[6])
				llego, _ := strconv.Atoi(res[7])

				fintentos := float64(intentos)
				fvalor := float64(valor)
				fllego := float64(llego)
				var impuesto float64
				impuesto = 0.3
				if prioridad == 1 {
					if strings.Compare(res[2], "pyme") == 0 {
						//Esta opcion nos dice que es la lista con prioridad por lo que multiplicamos el valor por si llego o no (1 si llego y 0 si no) y sumamos todo
						total = total + fvalor*fllego + fvalor*impuesto - fintentos
					} else {
						//Opcion de retail, da igual si llego, tenemos el dinero menos el nuemero de intentos
						total = total + fvalor - fintentos
					}
				} else {
					//Opcion pyme, valor*llego = Valor(0 o 1)
					total = total + fvalor*fllego - fintentos
				}
				Lista = append(Lista[:0], Lista[1:]...)
			}
			fmt.Println("----------------------------------")
			fmt.Println("Se tiene un total de: $", total, " hasta ahora.(Espere 3 segundo)")
			// 3 segundos para que el mensaje no se pierda muy rapido
			time.Sleep(time.Second)
			time.Sleep(time.Second)
			time.Sleep(time.Second)
		} else if strings.Compare(text, "3") == 0 {
			break
		} else {
			fmt.Println("Opcion introducida no existe")
		}
	}
}
