package controllers

import (
	"net/http"
)

func HealthCheck(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Write([]byte("OK!"))
}
