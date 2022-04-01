package main

import (
	"log"
	"net/http"
)

func home(rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    http.NotFound(rw, r)
    return
  }
  
  rw.Write([]byte("Bem vindo ao SnippetBox"))
}
func showSnippet(rw http.ResponseWriter, r *http.Request){
  rw.Write([]byte("Mostrar detalhes do Snippet"))
}
//curl -i -X POST http://localhost:4000/snippet/create

func createSnippet(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    rw.Header().Set("Allow","POST")
    http.Error(rw, "Método Não Permitido", http.StatusMethodNotAllowed)
    return
  }
  
  rw.Write([]byte("Criar um novo Snippet"))
}
func main() {
	mux := http.NewServeMux()
  mux.HandleFunc("/",home)
  mux.HandleFunc("/snippet",showSnippet)
  mux.HandleFunc("/snippet/create",createSnippet)
  log.Println("Iniciando o Servidor na Porta 4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
