package controllers

import (
	utiles "go_ecommerce/Utiles/Responses"
	"net/http"
)

func GetUserAddress( w http.ResponseWriter, r *http.Request ) {
	data := map[string]string{"health":"ok"}
	message := "health status is ok"
	response := utiles.Success(data, &message)

	response.Send( w )
}