package main

import (
	"main/database"
	"main/routers"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	router := mux.NewRouter().StrictSlash(true)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	routers.InitializeRouters(router, port)

}
