package routes

import (
	"github.com/dotsehyde/controller"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.Handle("/", controller.Home()).Methods("GET")
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controller.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controller.GetSingleBook).Methods("GET")
	router.HandleFunc("/book/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controller.DeleteBook).Methods("DELETE")
}
