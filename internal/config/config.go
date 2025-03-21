package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Log    LogConfig    `mapstructure:"Log"`
}

var globalConfig *Config

func Init() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")            //配置文件名称
	v.SetConfigType("yaml")              //配置文件类型
	v.AddConfigPath("./internal/config") //配置文件路径

	v.SetEnvPrefix("TASKMASTER") //环境变量前缀
	v.AutomaticEnv()             //自动绑定所有层级的环境变量

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败:%v", err)
	}

	//配置热更新
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("检测到配置变更", e.Name)
	})

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("配置解析失败:%v", err)
	}
	globalConfig = &cfg
	return globalConfig, nil
}

func GetConfig() *Config {
	return globalConfig
}
