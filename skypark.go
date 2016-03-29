package main

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	//"github.com/jinzhu/gorm"
)

type aa struct {
	name string
}

func welcome(c web.C, w http.ResponseWriter, r *http.Request) {
	t := template.New("welcome")
	t, _ = t.ParseFiles("views/welcome.html")
	a := aa {name: "qwe"}
	t.Execute(w, a)
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func main() {
	goji.Get("/", welcome)
	goji.Get("/hello/:name", hello)
	goji.Serve()
}