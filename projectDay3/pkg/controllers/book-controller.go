package controllers

import (
	"net/http"
	"github.com/jishnu21/projectday3/pkg/models"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"github.com/jishnu21/projectday3/pkg/utils"
	"strconv"
)


var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newbooks := models.GetAllBooks()
	res,_ := json.Marshal(newbooks)
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(res)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book ID")
	}
	bookDetails, _ := models.GetBookById(bookId)
	res, _ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(res)
	w.Write(res)
}

func Createbook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	newbook := &models.Book{}
	utils.ParseBody(r,newbook)
	b := newbook.CreateBook()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(res)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book ID")
	}
	delbook := models.DeleteBook(bookId)
	res,_ := json.Marshal(delbook)
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(res)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	updatebook := &models.Book{}
	utils.ParseBody(r,updatebook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id,err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("Error while parsing")
	}

	bookDetails,db:= models.GetBookById(Id)
	if updatebook.Title != ""{
		bookDetails.Title = updatebook.Title
	}
	if updatebook.Author != ""{
		bookDetails.Author = updatebook.Author
	}
	if updatebook.Year != 0 {
		bookDetails.Year = updatebook.Year
	}

	db.Save(&bookDetails)

	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(res)
	w.Write(res)
}