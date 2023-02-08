package handler

import (
	"net/http"

	"github.com/PGNahuel/FinHelp/internal/service"
	"github.com/gin-gonic/gin"
)

type deudaHandler struct {
	deudaService service.Deuda
}

func NewHandler(ds service.Deuda) *deudaHandler {
	return &deudaHandler{
		deudaService: ds,
	}
}

func (dh deudaHandler) ObtenerDeuda(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, "No implementado")
}
