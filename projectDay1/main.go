package main

import(
	"fmt"
	"net/http"
	"log"
)

func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 Not method",http.StatusNotFound)
	}

	if r.Method != "GET"{
		http.Error(w,"405 Method Not Allowed",http.StatusMethodNotAllowed)
	}

	fmt.Fprint(w,"Hello!")
}


func formhandler(w http.ResponseWriter,r *http.Request){
	if err := r.ParseForm(); err != nil {
		http.Error(w,"400 Bad Request",http.StatusBadRequest)
		return
	}

	if r.URL.Path != "/form"{
		http.Error(w,"404 Not Found",http.StatusNotFound)
	}

	if r.Method != "POST"{
		http.Error(w,"405 Method Not Allowed",http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Fprintf(w,"Name: %s\nEmail: %s\n", name, email)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formhandler)
	http.HandleFunc("/hello",helloHandler)


	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080",nil)
	 err != nil{
		log.Fatal(err)
	}

}