package server

import (
	"github.com/ayushjnv1/book-Crud-Op/book"
	"github.com/gorilla/mux"
)

var RegisterBookStore = func (router *mux.Router){
	router.HandleFunc("/book/",book.CreateBook).Methods("POST")
	router.HandleFunc("/book/",book.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",book.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}",book.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",book.DeleteBook).Methods(("DELETE"))
}