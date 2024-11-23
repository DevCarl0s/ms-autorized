package controladores

import (
	"genesis/ho/ms-autorized/comunes/contenedor"
	comunes_entidades "genesis/ho/ms-autorized/comunes/dominio/entidades"
	entidades_autorizacion "genesis/ho/ms-autorized/dominio/entidades/AutorizacionRecaudo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AutorizacionRecaudo(ctx *gin.Context) {

	peticion := &entidades_autorizacion.PeticionAutorizacionRecaudo{}
	err := ctx.ShouldBind(&peticion)
	if err != nil {
		errorMensaje := &comunes_entidades.RespuestaAutorizacion{
			Mensaje: err.Error(),
			Estado:  strconv.Itoa(http.StatusBadRequest),
		}

		ctx.JSON(http.StatusBadRequest, errorMensaje)
		return
	}
	respuesta, err := contenedor.ProcesarAutorizacion.Ejecutar(peticion)
	if err != nil {
		errorMensaje := &comunes_entidades.RespuestaAutorizacion{
			Mensaje: err.Error(),
			Estado:  "ERROR",
		}

		ctx.JSON(http.StatusInternalServerError, errorMensaje)
		return
	}

	ctx.JSON(respuesta.GetHttpEstado(), respuesta)

}
