package configs
import (
	"log"

	"github.com/spf13/viper"
)

// main file that setups all configs
func Setup() {
	
	// Set the names of the configuration files and the paths to search for them.
	viper.SetConfigName( "app" )
	viper.AddConfigPath( "./configs/yamls" )
	viper.AddConfigPath( "./configs" )
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf( "Error reading config file: %s", err )
	}
	

	viper.SetConfigName( "db" )
	viper.AddConfigPath( "./configs" )



	// Set up the environment variable prefix for Viper to search for
	viper.SetEnvPrefix( "APP" )
	viper.New().AutomaticEnv()
}
