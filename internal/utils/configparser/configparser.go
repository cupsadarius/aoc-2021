package configparser

import (
	"aoc/internal/utils/errors"
	"io/fs"
	"strings"

	"github.com/goccy/go-reflect"
	"github.com/spf13/viper"
)

// Parse configuration from the file and set the results to target.
// Allows the setting of default values through the defaults arg.
func Parse(file string, target interface{}, defaults ...map[string]interface{}) error {
	if reflect.ValueOf(target).Kind() != reflect.Ptr {
		return errors.ErrTargetNotPointer
	}

	viper.Reset()
	viper.SetConfigType("yaml")

	for _, def := range defaults {
		for k, v := range def {
			viper.SetDefault(k, v)
		}
	}

	if file != "" {
		viper.SetConfigFile(file)
	} else {
		viper.SetConfigName(".config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath("/etc/app/")
	}

	// https://github.com/spf13/viper/issues/584#issuecomment-451554896
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(*fs.PathError); ok && file != "" {
			return errors.ErrFileDoesNotExist
		}

		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	return viper.Unmarshal(target)
}
