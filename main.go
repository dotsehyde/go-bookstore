package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dotsehyde/routes"
	"github.com/gorilla/mux"
)

func main() {
	// file := http.FileServer(http.Dir("./static"))
	r := mux.NewRouter()
	//[Append] register routes
	// r.Handle("/", file).Methods("GET")
	routes.RegisterBookStoreRoutes(r)
	fmt.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
