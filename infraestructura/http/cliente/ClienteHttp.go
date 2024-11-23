package infraestructura_http_cliente

import (
	"bytes"
	comunes_entidades "genesis/ho/ms-autorized/comunes/dominio/entidades"
	"io"
	"log"
	"net/http"
	"time"
)

type ClienteHttp struct {
	cliente *http.Client
}

func (CHttp *ClienteHttp) Enviar(mensaje *comunes_entidades.HttpRequest) (*comunes_entidades.HttpResponse, error) {

	request, err := http.NewRequest(mensaje.Metodo, mensaje.ObtenerUrl(), bytes.NewBuffer(mensaje.Body))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	log.Printf("PETICION \n %+v", request)
	response, err := CHttp.cliente.Do(request)

	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	responesServicio := &comunes_entidades.HttpResponse{
		StatusCode: response.StatusCode,
		Body:       bodyBytes,
		Status:     response.Status,
	}

	//log.Printf("RESPUESTA RECIVIDA \n %+v", responesServicio)
	return responesServicio, nil
}

func InicializarCliente() (*ClienteHttp, error) {
	httpCliente := &http.Client{
		Timeout: time.Second * 5,
	}
	return &ClienteHttp{
		cliente: httpCliente,
	}, nil
}
