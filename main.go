// main.go
package main

import (
	"html/template"
	"net/http"
)

func main() {
	count := 0
    t, _ := template.ParseFiles("index.html")

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		t.Execute(w, count)
	})

	http.HandleFunc("/increment", func(w http.ResponseWriter, _ *http.Request) {
		count++
		t.Execute(w, count)
	})

	http.ListenAndServe(":8080", nil)
}

