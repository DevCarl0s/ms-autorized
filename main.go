package main

import (
	"fmt"
	"genesis/ho/ms-autorized/comunes/contenedor"
	adaptador_servidor "genesis/ho/ms-autorized/presentacion/config/adaptador"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("HO AUTORIZATION RECAUDADO")
	godotenv.Load()
	err := contenedor.InicializarContenedor()
	if err != nil {
		log.Panicf(err.Error())
	}
	log.SetOutput(os.Stdout)
	adaptador_servidor.Start()
}
