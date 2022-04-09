package routes

import (
	"github.com/DeeAmps/go-bookstore-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
