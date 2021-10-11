package main

import (
	"fmt"
	"log"
	"main/database"
	"main/routers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (app *App) initialiseRoutes() {
	app.Router = mux.NewRouter().StrictSlash(true)
	routers.InitializeRouters(app.Router)
}

func (app *App) run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), app.Router))
}

func (app *App) connectDatabase() {
	database.Connect()
}
