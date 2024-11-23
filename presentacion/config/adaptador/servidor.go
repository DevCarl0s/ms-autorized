package adaptador_servidor

import (
	"fmt"
	"genesis/ho/ms-autorized/presentacion/config/env"
	config "genesis/ho/ms-autorized/presentacion/config/routes"

	"log"
)

func Start() error {
	servidor, err := config.GinConfig()
	if err != nil {
		log.Fatal(err)
		return err
	}

	configEnv, err := env.LoadConfig()
	if err != nil {
		log.Fatal("Error cargando configuración:", err)
	}

	ip := configEnv.HOST_IP
	port := configEnv.HOST_PORT
	host := ip + ":" + port

	if configEnv.USE_HTTPS {
		fmt.Println("Iniciando servidor HTTPS en:", host)
		err = servidor.RunTLS(host, configEnv.SSL_CERT_FILE, configEnv.SSL_KEY_FILE)
		if err != nil {
			log.Fatal("Error iniciando servidor HTTPS:", err)
			return err
		}
	} else {
		fmt.Println("Iniciando servidor HTTP en:", host)
		err = servidor.Run(host)
		if err != nil {
			log.Fatal("Error iniciando servidor HTTP:", err)
			return err
		}
	}

	return nil

}
