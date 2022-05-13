package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/rmcs87/cc6m/pkg/models"
)

func (app *application) home(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(rw)
		return
	}
	files := []string{
		"./ui/html/home.page.tmpl.html",
		"./ui/html/base.layout.tmpl.html",
		"./ui/html/footer.partial.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(rw, err)
		return
	}
	err = ts.Execute(rw, nil)
	if err != nil {
		app.serverError(rw, err)
		return
	}
}

func (app *application) showSnippet(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(rw)
		return
	}

  s, err := app.snippets.Get(id)
  if err == models.ErrNoRecrod{
    app.notFound(rw)
    return
  }else if err != nil{
    app.serverError(rw, err)
    return
  }
  fmt.Fprintf(rw, "%v", s)
}

//curl -i -X POST http://localhost:4000/snippet/create
func (app *application) createSnippet(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		rw.Header().Set("Allow", "POST")
		app.clientError(rw, http.StatusMethodNotAllowed)
		return
	}
	//fake!!!!
	title := "10/05 aula"
	content := "COnteudo da aula de hoje: BD"
	expires := "30"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(rw, err)
	}
	http.Redirect(rw, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
