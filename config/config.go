package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"log"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	if err := c.initConfig(); err != nil {
		return nil
	}
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		// Parses the specified configuration file if it is specified.
		viper.SetConfigFile(c.Name)
	} else {
		// Default configuration file
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("LONGO")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// Monitor configuration file changes and hot loader
func (c *Config) watchConfig() {
	viper.WatchConfig()
	// Print log
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed:%s", e.Name)
	})
}
