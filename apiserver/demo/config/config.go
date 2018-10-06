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
		cfg,
	}

	if err := c.initConfig(); err != nil {
		return err
	}

	// 在初始化配置信息之后，再初始化日志
	c.initLog()

	c.watchConfig()

	return nil
}

func (c *Config) initLog() {
	// 构建日志信息
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则使用指定的配置文件
	} else {
		viper.AddConfigPath("demo/conf")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()            // 读取环境变量
	viper.SetEnvPrefix("APISERVER") //指定环境变量前缀
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil { // 解析配置文件
		return err
	}
	return nil
}

// 监控配置文件并热加载
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file change: %s", e.Name)
	})
}
