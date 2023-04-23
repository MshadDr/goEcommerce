package main

import (
	"fmt"
	"go_ecommerce/Routes"
	"go_ecommerce/configs"
	"log"
	"net/http"

	"github.com/spf13/viper"
)


func main() {
	// read all config files
	configs.Setup()

	appPort := viper.GetString( "app.port" )

	// load routes
	router := routes.NewRouter()

	// keep application up
	go func ()  {
		if err := http.ListenAndServe( ":" + appPort, router ); err != nil {
			log.Fatalf("Error starting server: %s", err)
		}
	}()
	
	fmt.Printf("Server is listening on port %s...\n", appPort)
	// blocks app indefinitely and never exits. 
	select {}
}
