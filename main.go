package main

import (
	"go_ecommerce/Routes"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	http.ListenAndServe( ":8098", router )
}