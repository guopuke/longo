package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/guopuke/longo/config"
	"github.com/guopuke/longo/router"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "Longo api server config file path.")
)

func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

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

	log.Printf("Start to lisenting the incoming requests on http address: %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
// 服务启动进行健康状态自检
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, e := http.Get(viper.GetString("url") + "/sd/health")
		if e == nil && resp.StatusCode == http.StatusOK {
			return nil
		}

		log.Print("Wating for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.\n")
}
