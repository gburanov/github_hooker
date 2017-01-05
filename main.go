package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Context struct {
	App *App
	Req *http.Request
	Res http.ResponseWriter
}

type CntFunc func(*Context) error

func wrap(app *App, mid CntFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		mid(&Context{
			App: app,
			Req: r,
			Res: w,
		})
	}
}

func main() {
	a := &App{}
	r := mux.NewRouter()

	get := r.Methods("GET").Subrouter()
	get.HandleFunc("/", wrap(a, statusHandler))

	post := r.Methods("POST").Subrouter()
	post.HandleFunc("/", wrap(a, webhookHandler))

	http.Handle("/", r)

	log.Println("Listening port 8080")
	http.ListenAndServe(":8080", nil)
}
