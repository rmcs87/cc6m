package main

import(
  "strconv"
  "fmt"
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
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1{
    http.NotFound(rw, r)
    return
  }
  fmt.Fprintf(rw, "Vais ser exibido o snippet de ID:%d", id)
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