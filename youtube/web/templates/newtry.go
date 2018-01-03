package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":9000", nil)

}
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "templates/index.gohtml", nil)

}
