package config

import (
	"genesis/ho/ms-autorized/presentacion/config/middleware"
	"genesis/ho/ms-autorized/presentacion/controladores"

	"github.com/gin-gonic/gin"
)

func GinConfig() (*gin.Engine, error) {
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.ValidarBody())
	router.POST("/api/autorizar-medio", controladores.AutorizacionRecaudo)
	return router, nil

}
