package controllers

import (
	"net/http"
	"go_ecommerce/Utiles/Responses"
)

func Health( w http.ResponseWriter, r *http.Request ) {

	data := map[string]string{"health":"ok"}
	message := "health status is ok"
	response := utiles.Success(data, &message)

	response.Send( w )
}