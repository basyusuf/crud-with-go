package main

import (
	"main/database"
	"main/routers"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	router := mux.NewRouter().StrictSlash(true)
	routers.InitializeRouters(router)
}
