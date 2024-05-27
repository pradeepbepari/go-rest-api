package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradeep/go-reat-api/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RoutesHandular(router)
	fmt.Println("listening port on : 8000")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
	// log.Fatal(http.ListenAndServe("localhost:8000", router))

}
