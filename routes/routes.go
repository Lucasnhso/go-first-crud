package routes

import (
	"crud/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", controllers.FindUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controllers.FindUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods(http.MethodDelete)


	return router
}