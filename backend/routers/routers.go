package routers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouters(router *mux.Router, port string) {
	UserRouters(router)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
