package helper

import "github.com/spf13/viper"

func GetViperEnvVariable(key string) string {

	value, ok := viper.Get(key).(string)

	if !ok {
		DefaultLoggingError("Key " + key + " not found in env file")
		return ""
	}

	return value
}