package routers

import (
	"main/controllers"

	"github.com/gorilla/mux"
)

func UserRouters(router *mux.Router) {
	routerPrefix := "/users"
	router.HandleFunc(routerPrefix, controllers.GetAllUser).Methods("GET")
	router.HandleFunc(routerPrefix+"/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc(routerPrefix, controllers.CreateUser).Methods("POST")
	router.HandleFunc(routerPrefix+"/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc(routerPrefix+"/{id}", controllers.DeleteUser).Methods("DELETE")
}
