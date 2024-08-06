package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/toledompm/my-monorepo/backend/domain/color"
	"github.com/toledompm/my-monorepo/backend/domain/healthcheck"
)

var chosenColor string

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/color", color.GetColorHandler(chosenColor)).Methods("GET")
	r.HandleFunc("/health", healthcheck.GetHealthcheckHandler()).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Server listening on port 8000")
	http.ListenAndServe(":8000", handlers.CORS()(r))
}
