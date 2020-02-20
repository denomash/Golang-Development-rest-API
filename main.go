package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article model
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

// Articles array
type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Endpoint HIT: Homepage")
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{Article{Title: "Welcome", Description: "Welcome Description", Content: "Hello world"}}

	fmt.Println("Endpoint HIT: Get all articles")
	json.NewEncoder(w).Encode(articles)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", getArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
