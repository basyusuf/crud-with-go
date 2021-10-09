package routers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouters(router *mux.Router) {
	UserRouters(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
