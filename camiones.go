package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
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

//Camion es el struct para los camiones donde se tiene un int que no dice si la orden se esta usando o no y un string reporte donde se guardan los reportes de las entregas
type Camion struct {
	Orden1   Orden
	Orden2   Orden
	usando1  int
	usando2  int
	Reporte1 string
	Reporte2 string
}

//obtenerOrden toma la primera orden de la lista y lo devuelve
func obtenerOrden(Lista []Orden) Orden {
	envio := Lista[0]
	return envio
}

//removerOrden remueve la primera orden de la lista y devuelve la lista
func removerOrden(Lista []Orden) []Orden {
	Lista = append(Lista[:0], Lista[1:]...)
	return Lista
}

//enviarCamion recibe un camion, trabaja con las ordenes que obtiene, enviandolas y escribe un reporte del envio ya haya sido recibido o no.
func enviarCamion(camion Camion) Camion {
	var Aux Orden
	var Aux2 Orden
	intentos1 := 0
	intentos2 := 0
	Listo1 := 0
	Listo2 := 0
	Aux = camion.Orden1
	valor1, _ := strconv.Atoi(Aux.Valor)
	valor2 := 0
	flag := 0

	if camion.usando2 != 0 {
		Aux2 = camion.Orden2
		valor2, _ = strconv.Atoi(Aux2.Valor)
		flag = 1
	}
	if flag == 1 {
		if valor1 > valor2 {
			if strings.Compare(Aux.Tienda, "pyme") == 0 {
				intentos1++
				if rand.Intn(100) < 81 {
					Listo1 = 1
				}
				intentos2++
				if rand.Intn(100) < 81 {
					Listo2 = 1

				}
				if Listo1 == 0 && valor1 > 10 {
					intentos1++
					if rand.Intn(100) < 81 {
						Listo1 = 1
					}
				}
				if Listo2 == 0 && valor2 > 10 {
					intentos2++
					if rand.Intn(100) < 81 {
						Listo2 = 1
					}
				}
			} else {
				intentos1++
				if rand.Intn(100) < 81 {
					Listo1 = 1
				}
				intentos2++
				if rand.Intn(100) < 81 {
					Listo2 = 1

				}
				if Listo1 == 0 {
					intentos1++
					if rand.Intn(100) < 81 {
						Listo1 = 1
					}
				}
				if Listo2 == 0 {
					intentos2++
					if rand.Intn(100) < 81 {
						Listo2 = 1
					}
				}
				if Listo1 == 0 {
					intentos1++
					if rand.Intn(100) < 81 {
						Listo1 = 1
					}
				}
				if Listo2 == 0 {
					intentos2++
					if rand.Intn(100) < 81 {
						Listo2 = 1
					}
				}

			}
		} else {
			if strings.Compare(Aux.Tienda, "pyme") == 0 {
				intentos2++
				if rand.Intn(100) < 81 {
					Listo2 = 1
				}
				intentos1++
				if rand.Intn(100) < 81 {
					Listo1 = 1
				}
				if Listo2 == 0 && valor2 > 10 {
					intentos2++
					if rand.Intn(100) < 81 {
						Listo2 = 1
					}
				}
				if Listo1 == 0 && valor1 > 10 {
					intentos1++
					if rand.Intn(100) < 81 {
						Listo1 = 1
					}
				}
			} else {
				intentos2++
				if rand.Intn(100) < 81 {
					Listo2 = 1
				}
				intentos1++
				if rand.Intn(100) < 81 {
					Listo1 = 1
				}
				if Listo2 == 0 && valor2 > 10 {
					intentos2++
					if rand.Intn(100) < 81 {
						Listo2 = 1
					}
				}
				if Listo1 == 0 && valor1 > 10 {
					intentos1++
					if rand.Intn(100) < 81 {
						Listo1 = 1
					}
				}
				if Listo2 == 0 && valor2 > 10 {
					intentos2++
					if rand.Intn(100) < 81 {
						Listo2 = 1
					}
				}
				if Listo1 == 0 && valor1 > 10 {
					intentos1++
					if rand.Intn(100) < 81 {
						Listo1 = 1
					}
				}
			}
		}
	} else {
		if strings.Compare(Aux.Tienda, "pyme") == 0 {
			intentos1++
			if rand.Intn(100) < 81 {
				Listo1 = 1
			}
			if Listo1 == 0 && valor1 > 10 {
				intentos1++
				if rand.Intn(100) < 81 {
					Listo1 = 1
				}
			}
		} else {
			intentos1++
			if rand.Intn(100) < 81 {
				Listo1 = 1
			}
			if Listo1 == 0 && valor1 > 10 {
				intentos1++
				if rand.Intn(100) < 81 {
					Listo1 = 1
				}
			}
			if Listo1 == 0 && valor1 > 10 {
				intentos1++
				if rand.Intn(100) < 81 {
					Listo1 = 1
				}
			}
		}
	}
	if flag == 1 {
		t := time.Now()
		timestamp := fmt.Sprintf("%02d-%02d-%d %02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
		camion.Reporte1 = timestamp + "@" + Aux.ID + "@" + Aux.Tienda + "@" + Aux.Prioridad + "@" + Aux.Destino + "@" + strconv.Itoa(intentos1) + "@" + Aux.Valor + "@" + strconv.Itoa(Listo1)
		camion.Reporte2 = timestamp + "@" + Aux2.ID + "@" + Aux2.Tienda + "@" + Aux2.Prioridad + "@" + Aux2.Destino + "@" + strconv.Itoa(intentos2) + "@" + Aux2.Valor + "@" + strconv.Itoa(Listo2)
	} else {
		t := time.Now()
		timestamp := fmt.Sprintf("%02d-%02d-%d %02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
		camion.Reporte1 = timestamp + "@" + Aux.ID + "@" + Aux.Tienda + "@" + Aux.Prioridad + "@" + Aux.Destino + "@" + strconv.Itoa(intentos1) + "@" + Aux.Valor + "@" + strconv.Itoa(Listo1)

	}
	return camion
}

func main() {
	var Aux Orden
	var Listapymes []Orden
	var Listaretail []Orden
	var Listaprio []Orden
	var conec *grpc.ClientConn
	conec, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conec.Close()

	c := chat.NewChatServiceClient(conec)
	var i int
	var CamP Camion
	CamP = Camion{
		usando1: 0,
		usando2: 0,
	}
	var CamR1 Camion
	CamR1 = Camion{
		usando1: 0,
		usando2: 0,
	}
	var CamR2 Camion
	CamR2 = Camion{
		usando1: 0,
		usando2: 0,
	}
	fmt.Println("----------------------------")
	fmt.Println("Los Camiones revisaran que no hayan nuevas entregas cada 7 segundos y partiran cada 35 segundos")
	var k int
	k = 0
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
		Aux = Orden{
			ID:        res1[0],
			Producto:  res1[1],
			Valor:     res1[2],
			Tienda:    res1[3],
			Destino:   res1[4],
			Prioridad: res1[5],
		}
		if strings.Compare(Aux.Prioridad, "Nada") == 0 {
			//No hay que hacer nada aca
		} else if strings.Compare(Aux.Prioridad, "1") == 0 {
			fmt.Println("Orden recibida:")
			fmt.Println(Aux)
			if strings.Compare(Aux.Tienda, "pyme") == 0 {
				Listaprio = append(Listaprio, Aux)
			} else {
				Listaretail = append(Listaretail, Aux)
			}
		} else {
			fmt.Println("Orden recibida:")
			fmt.Println(Aux)
			Listapymes = append(Listapymes, Aux)
		}

		if len(Listapymes)+len(Listaprio)+len(Listaretail) == 0 {
			fmt.Println("No hay entregas aun")
		} else {
			if len(Listaretail) != 0 {
				Aux = obtenerOrden(Listaretail)
				if CamR1.usando1 == 0 {
					//Reviso CamionRetail 1
					CamR1.Orden1 = Aux
					Listaretail = removerOrden(Listaretail)
					CamR1.usando1 = 1
				} else if CamR1.usando2 == 0 {
					CamR1.Orden2 = Aux
					Listaretail = removerOrden(Listaretail)
					CamR1.usando2 = 1
				} else if CamR2.usando1 == 0 {
					//Reviso CamionRetail2
					CamR2.Orden1 = Aux
					Listaretail = removerOrden(Listaretail)
					CamR2.usando1 = 1
				} else if CamR2.usando2 == 0 {
					CamR2.Orden2 = Aux
					Listaretail = removerOrden(Listaretail)
					CamR2.usando2 = 1
				} else {
					fmt.Println("Camiones de retails ocupados")
				}
			}
			if len(Listaprio) != 0 {
				Aux = obtenerOrden(Listaprio)
				if len(Listaretail) == 0 {
					if CamR1.usando1 == 0 {
						//Reviso CamionRetail 1
						CamR1.Orden1 = Aux
						Listaprio = removerOrden(Listaprio)
						CamR1.usando1 = 1
					} else if CamR1.usando2 == 0 {
						CamR1.Orden2 = Aux
						Listaprio = removerOrden(Listaprio)
						CamR1.usando2 = 1
					} else if CamR2.usando1 == 0 {
						//Reviso CamionRetail2
						CamR2.Orden1 = Aux
						Listaprio = removerOrden(Listaprio)
						CamR2.usando1 = 1
					} else if CamR2.usando2 == 0 {
						CamR2.Orden2 = Aux
						Listaprio = removerOrden(Listaprio)
						CamR2.usando2 = 1
					}
				} else {
					if CamP.usando1 == 0 {
						CamP.Orden1 = Aux
						Listaprio = removerOrden(Listaprio)
						CamP.usando1 = 1
					} else if CamP.usando2 == 0 {
						CamP.Orden2 = Aux
						Listaprio = removerOrden(Listaprio)
						CamP.usando2 = 1
					}
				}
			}
			if len(Listapymes) != 0 {
				if len(Listaprio) != 0 {
					fmt.Println("Orden en pyme con prioridad, Envio retrasado hasta que Cola prioridad termine")
				} else {
					Aux = obtenerOrden(Listapymes)
					if CamP.usando1 == 0 {
						CamP.Orden1 = Aux
						Listapymes = removerOrden(Listapymes)
						CamP.usando1 = 1
					} else if CamP.usando2 == 0 {
						CamP.Orden2 = Aux
						Listapymes = removerOrden(Listapymes)
						CamP.usando2 = 1
					}
				}

			}
		}
		if k != 5 {
			k++
		} else {
			k = 0
			if CamP.usando1 == 1 {
				CamP = enviarCamion(CamP)
				if CamP.usando2 == 1 {
					message := chat.Message{
						Body: CamP.Reporte2,
					}
					response, err := c.SayHola(context.Background(), &message)
					if err != nil {
						log.Fatalf("Error when calling server: %s", err)
					}
					fmt.Println(response.Body)
					CamP.usando2 = 0
				}
				message := chat.Message{
					Body: CamP.Reporte1,
				}
				response, err := c.SayHola(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling server: %s", err)
				}
				fmt.Println(response.Body)
				CamP.usando1 = 0
			}
			if CamR1.usando1 == 1 {
				CamR1 = enviarCamion(CamR1)
				//Reporte
				if CamR1.usando2 == 1 {
					message := chat.Message{
						Body: CamR1.Reporte2,
					}
					response, err := c.SayHola(context.Background(), &message)
					if err != nil {
						log.Fatalf("Error when calling server: %s", err)
					}
					fmt.Println(response.Body)
					CamR1.usando2 = 0
				}
				message := chat.Message{
					Body: CamR1.Reporte1,
				}
				response, err := c.SayHola(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling server: %s", err)
				}
				fmt.Println(response.Body)
				CamR1.usando1 = 0
				//
			}
			if CamR2.usando1 == 1 {
				CamR2 = enviarCamion(CamR2)
				//Reporte
				if CamR2.usando2 == 1 {
					message := chat.Message{
						Body: CamR2.Reporte2,
					}
					response, err := c.SayHola(context.Background(), &message)
					if err != nil {
						log.Fatalf("Error when calling server: %s", err)
					}
					fmt.Println(response.Body)
					CamR2.usando2 = 0
				}
				message := chat.Message{
					Body: CamR2.Reporte1,
				}
				response, err := c.SayHola(context.Background(), &message)
				if err != nil {
					log.Fatalf("Error when calling server: %s", err)
				}
				fmt.Println(response.Body)
				CamR2.usando1 = 0
				//
			}
		}
	}
}
