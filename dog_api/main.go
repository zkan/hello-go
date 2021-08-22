package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type breed struct{}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func GetBreeds(w http.ResponseWriter, r *http.Request) {
	url := "https://dog.ceo/api/breeds/list/all"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Ooops.. an error occurred, please try again")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("No response from request")
	}

	fmt.Fprint(w, string(body))

	// var bResp breed

	// if err := json.NewDecoder(resp.Body).Decode(&bResp); err != nil {
	// 	log.Fatal("ooopsss! an error occurred, please try again")
	// }

	// fmt.Println(bResp)
}

type RandomBreed struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func GetRandomBreed(w http.ResponseWriter, r *http.Request) {
	url := "https://dog.ceo/api/breeds/image/random"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Ooops.. an error occurred, please try again")
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)

	var breed RandomBreed
	if err := json.NewDecoder(resp.Body).Decode(&breed); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}

	fmt.Println(breed.Message)
	fmt.Println(breed.Status)

	json.NewEncoder(w).Encode(breed)
}

func main() {
	fmt.Println("Hello!")

	r := mux.NewRouter().StrictSlash(true)
	r.Use(commonMiddleware)

	r.HandleFunc("/breeds", GetBreeds).Methods("GET")
	r.HandleFunc("/random", GetRandomBreed).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
