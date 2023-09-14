package app

import (
	"github.com/shoelfikar/golang_backend_api/helper"
	"github.com/spf13/viper"
)

// ViperEnvVariable is func to get .env file
func NewViperLoad()  {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		// log.Println("Error while reading config file:", err)
		helper.DefaultLoggingError("Error while reading config file: " + err.Error())
		return
	}


	helper.DefaultLoggingInfo("success to load env config")
	
}
