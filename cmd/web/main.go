package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
  errorLog *log.Logger
  infoLog *log.Logger
}

func main() {
  addr := flag.String("addr", "4000", "Specifies the port for the server to startup on.");
  dsn := flag.String("dsn", "web:mdk00@web@/snippetbox?parseTime=true", "MySql data source name")

  flag.Parse();
  

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

  db, err := openDB(*dsn)
  if err != nil {
    errLog.Fatal(err)  
  }
  defer db.Close()

  app := application {
    errorLog: errLog,
    infoLog: infoLog,
  }

  server := &http.Server{
    Addr: ":"+*addr,
    ErrorLog: errLog,
    Handler: app.routes(),
  }
  infoLog.Printf("Server started on port %s\n", *addr);
  err = server.ListenAndServe()

  errLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
  db,err := sql.Open("mysql", dsn)
  if err != nil {
    return nil, err
  }
  if err = db.Ping(); err != nil {
    return nil, err
  }
  return db, nil
}
