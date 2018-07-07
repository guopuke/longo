package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"longo/router"
)

func main() {
	// Create the Gin engine
	g := gin.New()

	// Gin middlewares
	middlewares := []gin.HandlerFunc{}

	router.Load(
		// Cores
		g,

		// Middlewares
		middlewares...,
	)

	fmt.Printf("Start to lisenting the incoming requests on http address: %s", ":8080")
	fmt.Printf(http.ListenAndServe(":8080", g).Error())
}
