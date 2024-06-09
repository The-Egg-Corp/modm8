package core

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var (
	persistenceCfg = viper.New()
	settingsCfg    = viper.New()
)

func SetupConfig(instance viper.Viper, name string, fileType string) {
	instance.SetConfigName(name)
	instance.SetConfigType(fileType)
	instance.AddConfigPath(ConfigDir())
}

func WriteToConfig(instance viper.Viper, i interface{}) {
	var result map[string]interface{}
	err := mapstructure.Decode(i, &result)
	if err != nil {
		fmt.Printf("%#v", result)
		return
	}

	instance.MergeConfigMap(result)
}
