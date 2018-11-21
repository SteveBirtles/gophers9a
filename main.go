package main

import (
	"net/http"
	"encoding/json"
)

type Post struct {
	Name string `json:"name"`
	Drink string `json:"drink"`
}

var posts = make([]Post, 0)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "index.html")

}

func listHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)

}

func postHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	posts = append(posts, Post{Name: r.FormValue("name"), Drink: r.FormValue("drink")})

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/list", listHandler)

	http.ListenAndServe(":8081", http.DefaultServeMux)

}