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
  
  srv := &http.Server{
    Addr:     *addr,
		ErrorLog: errorLogger,
		Handler:  app.routes(),
  }
  
  infoLogger.Printf("Iniciando o Servidor na Porta %s\n", *addr)
  err := srv.ListenAndServe()
  errorLogger.Fatal(err)
}
