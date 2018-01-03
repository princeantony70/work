package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/test.html", TestHandler)
	http.ListenAndServe(":8080", nil)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	//Parsing HTML
	t, err := template.ParseFiles("test.html")
	if err != nil {
		fmt.Println(err)
	}

	items := struct {
		Name string
		City string
	}{
		Name: "MyName",
		City: "MyCity",
	}

	t.Execute(w, items)
}
