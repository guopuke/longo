package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/guopuke/longo/config"
	"github.com/guopuke/longo/model"
	"github.com/guopuke/longo/router"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"reflect"
	"sync"
	"time"
)

var (
	cfg                         = pflag.StringP("config", "c", "", "Longo api server config file path.")
	_   binding.StructValidator = &defaultValidator{}
)

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func main() {
	binding.Validator = new(defaultValidator)

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

func (v *defaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}

	return nil
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		// add any custom validations etc. here
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
