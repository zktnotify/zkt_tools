package initialize

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func SetupConfig() {
	var (
		conf = flag.String("conf", "./configs/config", "config file path")
	)
	flag.Parse()
	fmt.Printf("using config:%s\n", *conf)
	viper.SetConfigName(*conf)
	// 优先从configmap读取
	viper.AddConfigPath("/etc/config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Read config fail:%v\n", err.Error())
	}
}
