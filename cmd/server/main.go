package main

import (
	"github.com/PGNahuel/FinHelp/cmd/server/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.MapRouter(r)

	r.Run(":8080")
}
