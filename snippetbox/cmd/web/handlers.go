package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func(app *application) home(rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    app.notFound(rw)
    return
  }
  files := []string{
    "./ui/html/home.page.tmpl.html",
    "./ui/html/base.layout.tmpl.html",
    "./ui/html/footer.partial.tmpl.html",
  }
  
  ts, err := template.ParseFiles(files...)
  
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, nil)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) showSnippet(rw http.ResponseWriter, r *http.Request){
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1{
    app.notFound(rw)
    return
  }
  fmt.Fprintf(rw, "Vais ser exibido o snippet de ID:%d", id)
}
//curl -i -X POST http://localhost:4000/snippet/create

func(app *application) createSnippet(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    rw.Header().Set("Allow","POST")
    app.clientError(rw, http.StatusMethodNotAllowed)
    return
  }
  
  rw.Write([]byte("Criar um novo Snippet"))
}