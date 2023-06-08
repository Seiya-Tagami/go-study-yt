// @see https://www.youtube.com/watch?v=YMQUQ6XQgz8&list=PLzUGFf4GhXBL4GHXVcMMvzgtO8-WEJIoY&index=2

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title: "Test Title", 
			Desc: "Test Description", 
			Content: "Hello World"},
	}

	fmt.Println("Endpoint Hi: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

type Articles []Article

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}