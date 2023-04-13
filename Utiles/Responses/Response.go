package utiles

import (
	"encoding/json"
	"net/http"
)

type Response struct{
	Status		string			`json:"statusCode"`
	Message		*string			`json:"message,omitempty"`
	Code		int				`json:"code"`
	Data		interface{}	`json:"data,omitempty"`
}

func Success(data interface{}, message *string) Response{
	return Response{
		Status: "success",
		Message: message,
		Code: 200,
		Data: data,
	}
}

func (r Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	json.NewEncoder(w).Encode(r)
}
