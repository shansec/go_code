package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type DataBaseConnection struct {
	IPAddress    string
	Port         int
	UserName     string
	Password     int
	DataBaseName string
}

func main() {
	config := viper.New()
	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed", e.Name)
	})
	// 配置文件名称
	config.SetConfigName("config")
	config.AddConfigPath("E:\\Go\\study\\viper")
	// 设置文件类型
	config.SetConfigType("yaml")
	var dataBaseInfo DataBaseConnection

	err := config.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := config.Unmarshal(&dataBaseInfo); err != nil {
		panic(fmt.Errorf("read config file to struct err: %s \n", err))
	}

	fmt.Println(dataBaseInfo)
	fmt.Println(config.Get("IPAddress"))
}
