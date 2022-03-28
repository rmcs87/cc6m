package main

import (
	"log"
	"net/http"
)

func home(rw http.ResponseWriter, r *http.Request){
  rw.Write([]byte("Bem vindo ao SnippetBox"))
}

func main() {
	mux := http.NewServeMux()
  mux.HandleFunc("/",home)
  log.Println("Iniciando o Servidor na Porta 4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
