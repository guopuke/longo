package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guopuke/longo/config"
	"github.com/guopuke/longo/model"
	v "github.com/guopuke/longo/pkg/version"
	"github.com/guopuke/longo/router"
	"github.com/guopuke/longo/router/middleware"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
)

var (
	cfg     = pflag.StringP("config", "c", "", "Longo api server config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	pflag.Parse()

	if *version {
		vInfo := v.Get()

		marshal, err := json.MarshalIndent(&vInfo, "", " ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(marshal))
		return
	}

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
		middleware.Logging(),
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.\n", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	// Start to listening the incoming requests.
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming request on https address: %s", viper.GetString("tls.addr"))
			log.Infof(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}

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
