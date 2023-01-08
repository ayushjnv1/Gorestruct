package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartAPIServer(){
	r := mux.NewRouter()
	RegisterBookStore(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r));
}