package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

var Posts []Post

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! please visit /posts to view the json")
}

func posts(w http.ResponseWriter, r *http.Request) {
	Posts = []Post{
		{Title: "Tutorial: Developing a RESTful API with Go", Description: "some description", Content: "post contents here"},
		{Title: "Tutorial: Developing an HTTP Server", Description: "some description", Content: "post contents here"},
	}
	json.NewEncoder(w).Encode(Posts)
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/posts", posts)
	http.ListenAndServe(":3000", nil)
}

func main() {
	fmt.Println("Server started at http://localhost:3000")
	handleRequests()
}
