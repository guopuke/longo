package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guopuke/longo/router"
	"log"
	"net/http"
	"time"
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

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.\n", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	fmt.Printf("Start to lisenting the incoming requests on http address: %s\n", ":8080")
	fmt.Printf(http.ListenAndServe(":8080", g).Error())
}

// pingServer pings the http server to make sure the router is working.
// 服务启动进行健康状态自检
func pingServer() error {
	for i := 0; i < 2; i++ {
		resp, e := http.Get("http://localhost:8080" + "/sd/health")
		if e == nil && resp.StatusCode == http.StatusOK {
			return nil
		}

		log.Print("Wating for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.\n")
}
