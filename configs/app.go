package configs

import "github.com/spf13/viper"

type AppConfig struct {
	Port		string	`mapstructure:"port"`
	Environment	string	`mapstructure:"environment"`
}

// get Applications configs
func GetAppConfig() AppConfig {
	var config AppConfig
	err := viper.UnmarshalKey( "app", &config )
	if err != nil {
		panic( err )
	}
	return config
}
