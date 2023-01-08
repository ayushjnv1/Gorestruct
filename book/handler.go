package book

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ayushjnv1/book-Crud-Op/api"
	"github.com/ayushjnv1/book-Crud-Op/db"
	"github.com/ayushjnv1/book-Crud-Op/utils"
	"github.com/gorilla/mux"
)


var NewBook db.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := db.GetAllBooks()
	
	api.Success(w,http.StatusOK,newBooks)	
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    bookId := vars["bookId"]
	ID,err := strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Printf("error while parse")
	}
	bookDetails,_ := db.GetBookById(ID)
	api.Success(w,http.StatusOK,bookDetails)
}

func CreateBook(w http.ResponseWriter , r *http.Request){
	CreateBook := &db.Book{}
	utils.ParseBody(r,CreateBook)
	b:= CreateBook.CreateBook()
	api.Success(w,http.StatusOK,b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    bookId := vars["bookId"]
	ID,err := strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Printf("error while parse")
	}
	
	book:= db.DeleteBook(ID)
	api.Success(w,http.StatusOK,book)

}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &db.Book{}
	
	utils.ParseBody(r,updateBook)	
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID,err := strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Printf("error while parse")
	}
    bookDetails,db := db.GetBookById(ID)
	if updateBook.Name!=""{
		bookDetails.Name =  updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	api.Success(w,http.StatusOK,bookDetails)

}