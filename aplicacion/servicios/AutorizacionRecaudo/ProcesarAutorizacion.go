package servicio_autorizacionrecaudo

import (
	casosuso_comunes "genesis/ho/ms-autorized/comunes/aplicacion/casosuso"
	"genesis/ho/ms-autorized/comunes/dominio/constantes"
	comunes_entidades "genesis/ho/ms-autorized/comunes/dominio/entidades"
	entidades_autorizacion "genesis/ho/ms-autorized/dominio/entidades/AutorizacionRecaudo"
	"log"
	"net/http"
)

type ProcesarAutorizacion struct {
	InformacionCliente *casosuso_comunes.ObtenerInformacionCliente
}

func (PIC *ProcesarAutorizacion) Ejecutar(peticion *entidades_autorizacion.PeticionAutorizacionRecaudo) (*comunes_entidades.RespuestaAutorizacion, error) {
	log.Println("Ejecutando servicio de autorizacion de recaudo")
	cliente, _ := PIC.InformacionCliente.Ejecutar(peticion)

	respuesta := &comunes_entidades.RespuestaAutorizacion{}
	statusCode := http.StatusBadRequest
	status := constantes.NO_AUTORIZADO
	var data *entidades_autorizacion.ClienteData

	if cliente.Success && len(cliente.Data) > 0 {
		statusCode = http.StatusOK
		status = constantes.AUTORIZADO
		data = &cliente.Data[0]
	}

	respuesta.Data = data
	respuesta.Mensaje = cliente.Message
	respuesta.Estado = status
	respuesta.SetHttpEstado(statusCode)

	return respuesta, nil
}
