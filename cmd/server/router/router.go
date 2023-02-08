package router

import (
	"github.com/PGNahuel/FinHelp/cmd/server/handler"
	"github.com/gin-gonic/gin"
)

func MapRouter(server *gin.Engine) {
	groupRoutes := server.Group("/api")
	{
		buildDeudaRoutes(groupRoutes)
	}
}

func buildDeudaRoutes(deudaRoutes *gin.RouterGroup) {
	dh := handler.NewHandler(nil)

	dr := deudaRoutes.Group("/deuda")
	{
		dr.GET("/", dh.ObtenerDeuda)
	}
}
