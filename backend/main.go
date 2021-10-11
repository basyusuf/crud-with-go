package main

import (
	"main/database"
	"main/routers"
	"os"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func main() {
	database.Connect()
	Router = mux.NewRouter().StrictSlash(true)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	routers.InitializeRouters(Router, port)

}
