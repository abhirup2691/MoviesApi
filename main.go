package main

import (
	"encoding/json" //for converting our data back into json
	"fmt"
	"log"
	"math/rand" // to generate random nos
	"net/http"  // to create a server
	"strconv"   // for string conversion

	"github.com/gorilla/mux" // to use gorilla mux that we have downloaded // gorilla mux is used for routing
)

type Movie struct{

	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`  //  this way we are defining because when data comes from Post man, it wil be easier to encode and decode the data

}

type Director struct {

	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie

func main(){
	
	movies = append(movies, Movie{ID: "1",Isbn: "43567",Title: "Return to Madagasgar",Director: &Director{FirstName: "John",LastName: "Massarati"} })
	movies = append(movies, Movie{ID: "2",Isbn: "43553",Title: "Koi Mil Gaya",Director: &Director{FirstName: "Rajesh",LastName: "Roshan"} })
	r:= mux.NewRouter()

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at Port 8000")
	log.Fatal(http.ListenAndServe(":8000",r)) //http.ListenAndServe(":8000",r) starts the server at port 8000
}

func getMovies(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json") //  setting the type we are going to return back to postman
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r) // Fetching the request parameters mux.Vars(r) will hold all the request parameters

	for index,item := range(movies){
		if item.ID == params["id"] {
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}

}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)

	for _, item := range(movies){
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
			
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)
	for index,item := range(movies){
		if item.ID == params["id"] {
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = params["id"]
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}