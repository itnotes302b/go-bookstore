package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/itnotes302b/go-bookstore/pkg/utils"
	"github.com/itnotes302b/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter , r *http.Request)  {
	NewBooks := models.GetAllBooks()
	res , _ := json.Marshal(NewBooks)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter , r *http.Request)  {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID , err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing ")
	}
	bookDetails,_ := models.GetBookById()
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
}

func CreateBook(w http.ResponseWriter, r *http.Request)  {
	CreateBook := &models.Book{}
	utils.ParseBody(r,CreateBook)
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter,r *http.Request)  {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID ,err := strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res , _ := json.Marshal(book)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter,r *http.Request)  {
	var UpdateBook = &models.Book{}
	utils.ParseBody(r, UpdateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID , err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing ")
	}
	bookDetails, db := models.GetBookById(ID)
	if UpdateBook.Name != "" {
		bookDetails.Name = UpdateBook.Name
	} 
	if UpdateBook.Author != "" {
		bookDetails.Author = bookDetails.Author
	}
	if UpdateBook.Publication != "" {
		bookDetails.Publication = UpdateBook.Publication
	}
	db.Save(&bookDetails)
	res ,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}