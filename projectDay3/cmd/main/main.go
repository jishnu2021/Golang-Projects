package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jishnu21/projectday3/pkg/routes"
	// "github.com/go-sql-driver/mysql"
)

func main(){
	r := mux.NewRouter();
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}