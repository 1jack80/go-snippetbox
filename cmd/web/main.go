package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
  addr := flag.String("addr", "4000", "Specifies the port for the server to startup on.");
  flag.Parse();

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

  mux := http.NewServeMux()
  fileserver := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/", http.StripPrefix("/static", fileserver))

  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet/view", snippetView)
  mux.HandleFunc("/snippet/create", snippetCreate)

  infoLog.Printf("Server started on port %s\n", *addr);

  server := &http.Server{
    Addr: ":"+*addr,
    ErrorLog: errLog,
    Handler: mux,
  }
  err := server.ListenAndServe()

  errLog.Fatal(err)
}
