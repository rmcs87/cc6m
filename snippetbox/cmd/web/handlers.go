package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    http.NotFound(rw, r)
    return
  }
  files := []string{
    "./ui/html/home.page.tmpl.html",
    "./ui/html/base.layout.tmpl.html",
    "./ui/html/footer.partial.tmpl.html",
  }
  
  ts, err := template.ParseFiles(files...)
  
  if err != nil{
    log.Println(err.Error())
    http.Error(rw, "Internal Server Error", 500)
    return
  }
  err = ts.Execute(rw, nil)
  if err != nil{
    log.Println(err.Error())
    http.Error(rw, "Internal Server Error", 500)
    return
  }
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