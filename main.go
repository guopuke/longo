package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/guopuke/longo/config"
	"github.com/guopuke/longo/model"
	"github.com/guopuke/longo/router"
	"github.com/guopuke/longo/router/middleware"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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

	// init DB
	model.DB.Init()
	defer model.DB.Close()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine
	g := gin.New()

	router.Load(
		// Cores
		g,

		// Middlewares
		middleware.RequestId(),
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.\n", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
// 服务启动进行健康状态自检
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, e := http.Get(viper.GetString("url") + "/sd/health")
		if e == nil && resp.StatusCode == http.StatusOK {
			return nil
		}

		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.\n")
}
