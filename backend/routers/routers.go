package routers

import (
	"main/controllers"

	"github.com/gorilla/mux"
)

func InitializeRouters(router *mux.Router) {
	UserRouters(router)
	router.HandleFunc("/", controllers.HealthCheck).Methods("GET")
}
