package main

import (
	"fmt"
	"log"	
	"net/http"
)

func greeting(text string) string {
	return "<b>" + text + "</b>"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, greeting("Code.education Rocks!"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}