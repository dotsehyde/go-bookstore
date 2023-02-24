package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	r := mux.NewRouter()
	r.Handle("/", fileServer).Methods("GET")
	fmt.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
