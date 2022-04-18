package main

import (
  "flag"
	"log"
	"net/http"
  "os"
)

type application struct{
  errorLog *log.Logger
  infoLog *log.Logger
}

func main() {
  addr := flag.String("addr", ":4000", "Porta do Endereço")
  flag.Parse()

  infoLogger := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
  errorLogger := log.New(os.Stderr, "ERRO:\t", log.Ldate|log.Ltime|log.Lshortfile)

  app := &application{
    errorLog: errorLogger, 
    infoLog: infoLogger,
  }
  
	mux := http.NewServeMux()
  mux.HandleFunc("/",app.home)
  mux.HandleFunc("/snippet",app.showSnippet)
  mux.HandleFunc("/snippet/create",app.createSnippet)

  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/",http.StripPrefix("/static",fileServer))
  
  infoLogger.Printf("Iniciando o Servidor na Porta %s\n", *addr)
  err := http.ListenAndServe(*addr, mux)
  errorLogger.Fatal(err)
}
