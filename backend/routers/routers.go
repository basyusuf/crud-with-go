package routers

import (
	"fmt"
	"log"
	"main/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouters(router *mux.Router, port string) {
	UserRouters(router)
	router.HandleFunc("/", controllers.HealthCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
