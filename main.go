package main

import (
	"log"
	"net/http"
)

// view all snippets
func home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    log.Println(r.URL)
    http.NotFound(w, r)
    return
  }

  w.Write([]byte("Hello world from snippetbox"))
}

// Create a new snippet
func snippetCreate(w http.ResponseWriter, r *http.Request) {
  
  if r.Method != "POST" {
    w.Header().Set("Allow", "POST")
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write([]byte("Method not allowed"))
    return
  }

  w.Write([]byte("Create a new snippet"))
}

// view a specific snippet
func snippetView(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("View a specific snippet"))
}

func main() {
  // create a new mux
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet/create", snippetCreate)
  mux.HandleFunc("/snippet/view", snippetView)

  // listen and serve using the mux created
  log.Println("Server started on port 4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
