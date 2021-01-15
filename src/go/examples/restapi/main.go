package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to home page!")
	fmt.Println("End point hit: homePage")

}

func handleRequests() {

	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))

}

func main() {
	handleRequests()
}
