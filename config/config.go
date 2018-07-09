package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"

	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return nil
	}

	// 初始化日志包
	c.initLog()

	// 监听配置文件并实现热更新
	c.watchConfig()

	return nil
}

func (c *Config) initLog() {
	PassLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&PassLagerCfg)
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
		log.Infof("Config file changed: %s", e.Name)
	})
}
