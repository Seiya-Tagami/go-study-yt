// @see https://www.youtube.com/watch?v=W5b64DXeP0o&list=PLzUGFf4GhXBL4GHXVcMMvzgtO8-WEJIoY

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}