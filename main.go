package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lautarooyuela/go-example-project/routes"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/healthy", routes.HealthyHandler)

	http.ListenAndServe(":3000", r)

}
