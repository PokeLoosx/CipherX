package initialize

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"CipherX/config"
)

// Viper Initialize configuration
func Viper(path string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
		if err = v.Unmarshal(&config.GinConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&config.GinConfig); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Viper initialized successfully")
	return v
}
