package appcore

import (
	"fmt"
	"modm8/backend/common/paths"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func SetupConfig(instance *viper.Viper, name string, fileType string) {
	instance.SetConfigName(name)
	instance.SetConfigType(fileType)
	instance.AddConfigPath(paths.ConfigDir())
}

func Save(instance *viper.Viper, i any, path string) error {
	// Write struct values to viper config
	WriteToConfig(instance, i)

	// Write config to the file at the specified path.
	if err := instance.WriteConfigAs(path); err != nil {
		return err
	}

	return nil
}

func WriteToConfig(instance *viper.Viper, i any) error {
	var result map[string]any
	err := mapstructure.Decode(i, &result)
	if err != nil {
		fmt.Printf("%#v", result)
		return err
	}

	return instance.MergeConfigMap(result)
}

func ReadOrCreate(instance *viper.Viper, i any, path string) error {
	if err := instance.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No file found, creating: " + path)
			return Save(instance, i, path)
		}

		return err
	}

	if err := instance.Unmarshal(i); err != nil {
		return err
	}

	return nil
}
