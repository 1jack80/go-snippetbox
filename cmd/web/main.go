package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
  addr := flag.Int("addr", 4000, "Specifies the port for the server to startup on.");
  flag.Parse();

  mux := http.NewServeMux()

  fileserver := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/", http.StripPrefix("/static", fileserver))

  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet/view", snippetView)
  mux.HandleFunc("/snippet/create", snippetCreate)

  fmt.Printf("Server started on port %d \n", addr);
  err := http.ListenAndServe(fmt.Sprintf(":%d",addr), mux)

  log.Fatal(err)
}
