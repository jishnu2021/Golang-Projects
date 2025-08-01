package main

import (
	"fmt" 
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"math/rand"
	"strconv"
)

type Movie struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Rating float64 `json:"rating"`
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie


func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getEachMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","Application/json")
	params := mux.Vars(r)

	for _,item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}


func createMovie(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","Application/json")
	var movie Movie
	 _ = json.NewDecoder(r.Body).Decode(&movie)
	 movie.ID = strconv.Itoa(rand.Intn(10000000))
	 movies = append(movies,movie)

	 json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","Application/json")
	params := mux.Vars(r)
	for index,data := range movies{
		if data.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
				var movie Movie
	 			_ = json.NewDecoder(r.Body).Decode(&movie)
	 			movie.ID = params["id"]
	 			movies = append(movies,movie)
	 			json.NewEncoder(w).Encode(movie)
				return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Title: "Inception", Year: 2010, Rating: 8.8, Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})
	movies = append(movies, Movie{ID: "2", Title: "The Matrix", Year: 1999, Rating: 8.7, Director: &Director{FirstName: "Lana", LastName: "Wachowski"}})
	movies = append(movies, Movie{ID: "3", Title: "Interstellar", Year: 2014, Rating: 8.6, Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getEachMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")


	fmt.Print("Starting the server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080",r))

}