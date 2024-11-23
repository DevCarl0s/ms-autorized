package comunes_entidades

import (
	"log"
	"regexp"
)

type HttpRequest struct {
	Body      []byte
	Url       string
	Metodo    string
	Protocolo string
}

func (HR *HttpRequest) ObtenerUrl() string {
	if HR.Protocolo == "" {
		HR.Protocolo = "http"
	}

	urlOk, err := regexp.Match(`^http?`, []byte(HR.Url))
	if err != nil {
		log.Println("NO MATCH POR ERROR", err)
		urlOk = false

	}

	if urlOk {
		return HR.Url
	}

	return HR.Protocolo + "://" + HR.Url

}
