package routers

import (
	"main/controllers"

	"github.com/gorilla/mux"
)

func UserRouters(router *mux.Router) {
	routerPrefix := "/users"
	router.HandleFunc(routerPrefix, controllers.GetAllUser).Methods("GET")
	router.HandleFunc(routerPrefix+"/{id:[0-9]+}", controllers.GetUserById).Methods("GET")
	router.HandleFunc(routerPrefix, controllers.CreateUser).Methods("POST")
	router.HandleFunc(routerPrefix+"/{id:[0-9]+}", controllers.UpdateUser).Methods("PATCH")
	router.HandleFunc(routerPrefix+"/{id:[0-9]+}", controllers.DeleteUser).Methods("DELETE")
}
