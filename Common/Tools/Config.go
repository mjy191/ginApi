package Tools

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath("./Config")
	err := Config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file:%s\n", err))
	}
}
