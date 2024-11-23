package middleware

import (
	"bytes"
	"errors"
	"fmt"
	comunes_entidades "genesis/ho/ms-autorized/comunes/dominio/entidades"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func ValidarBody() gin.HandlerFunc {
	return validarBodyVacio
}

func validarBodyVacio(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		mensajeError := &comunes_entidades.RespuestaAutorizacion{Estado: "JSON mal formado", Mensaje: err.Error()}
		log.Println("MENSAJE ERROR", mensajeError)
		ctx.JSON(http.StatusBadRequest, mensajeError)
		ctx.Abort()

		return
	}
	match, err := regexp.Match(`^\s*\{\s*\}\s*$`, bodyBytes)
	if err != nil {
		fmt.Print(err)
		mensajeError := &comunes_entidades.RespuestaAutorizacion{Estado: "JSON mal formado", Mensaje: err.Error()}
		fmt.Println("MENSAJE ERROR", mensajeError)
		ctx.JSON(http.StatusBadRequest, mensajeError)
		ctx.Abort()
		return
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	if len(bodyBytes) == 0 || match {
		err := errors.New("json vacio o malformado ")
		mensajeError := &comunes_entidades.RespuestaAutorizacion{Estado: "JSON mal formado", Mensaje: err.Error()}
		fmt.Println("MENSAJE ERROR", mensajeError)
		ctx.JSON(http.StatusBadRequest, mensajeError)
		ctx.Abort()
		return
	}

	ctx.Next()
}
