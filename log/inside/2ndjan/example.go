package main
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()


/*type FileHandler struct {
  Path string
}

func (f FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  prefix := "public"
  http.ServeFile(w, r, path.Join(prefix, f.Path))
}
*/

r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
vars := mux.Vars(r)
title := vars["title"]
page := vars["page"]

fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
})

http.ListenAndServe(":8080", r)
}
/*
func main() {
    var dir string

    flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
    flag.Parse()
    r := mux.NewRouter()

    // This will serve files under http://localhost:8000/static/<filename>
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

    srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8000",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
} */
