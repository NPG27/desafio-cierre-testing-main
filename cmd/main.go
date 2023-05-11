package main

import (
	"log"

	"github.com/bootcamp-go/desafio-cierre-testing/cmd/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	err := r.Run(":18085")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %s", err)
	}

}
