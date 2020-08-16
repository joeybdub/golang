package main

import (
	"fmt"
	"log"
	"net/http"
)

type Content struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./test.json"
	}
	http.ServeFile(w, r, p)

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8099", nil))
}

func main() {
	handleRequests()
}
