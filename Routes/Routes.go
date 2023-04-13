package routes

import (
	"go_ecommerce/Controllers"
	"net/http"
)

type Route struct{
	Name	string
	Method	string
	Pattern	string
	Handler	http.HandlerFunc
}

type Routes []Route // Routes is a slice of Route struct

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.Handler)
	}
	return router
}

var routes = Routes{
	Route{
		"checkHealth",
		"GET",
		"/health",
		controllers.Health,
	},
}


