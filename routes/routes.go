package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradeep/go-reat-api/controller"
)

var RoutesHandular = func(routes *mux.Router) {

	routes.HandleFunc("/register", controller.CreateUser).Methods(http.MethodPost)
	routes.HandleFunc("/users", controller.GetAllUser).Methods(http.MethodGet)
	routes.HandleFunc("/users/{user-id}", controller.GetUserbyId).Methods(http.MethodGet)
	routes.HandleFunc("/users/{user-id}", controller.UpdateUser).Methods(http.MethodPut)
	routes.HandleFunc("/users/{user-id}", controller.DeleteUser).Methods(http.MethodDelete)
}
