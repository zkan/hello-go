package main

import (
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

func main() {
	fmt.Println("Hello!")

	var router *mux.Router
	router = mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)

	router.HandleFunc("/breeds", GetBreeds).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
