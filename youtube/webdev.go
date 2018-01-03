package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello it is index ")

	v:=r.FormValue("q")

	w.Header().Set("Content-type","text/html; charset=utf-8")

	io.WriteString(w,'
	<form method ="POST">
	<input type="text" name ="q" id="q" value="Hello">
	<input type ="submit">
	</form>
	<br>' v)
}

func home_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is home  ")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/home/", home_handler)
	http.ListenAndServe(":8000", nil)
}
