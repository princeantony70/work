package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HELLO WORLD MMMMMM kkkk ")

}

func main() {

	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":9000", nil)
}
