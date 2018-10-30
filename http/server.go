package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/amsourav/learn-go/basics/validationkit"
)

type Gopher struct {
	Name string
}

func greet(w http.ResponseWriter, r *http.Request) {
	var gophername string

	gophername = r.URL.Query().Get("gophername")

	if gophername == "" {
		gophername = "Gopher"
	}

	gopher := Gopher{Name: gophername}
	renderTemplate(w, "./template/hello.html", gopher)
}

func renderTemplate(w http.ResponseWriter, templateFile string, templateData Gopher) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Encontered a error while parsing the template", err)
	}

	t.Execute(w, templateData)
}

func validUserNameHandler(w http.ResponseWriter, r *http.Request) {
	var userNameSyntaxResult bool

	username := r.URL.Query().Get("username")

	if username == "" {
		http.Error(w, "Username not provided", http.StatusInternalServerError)
	} else {
		userNameSyntaxResult = validationkit.ValidaterUsername(username)
		fmt.Fprintf(w, "Syntax check result for %v is %v", username, userNameSyntaxResult)
	}
}

func main() {
	// MakeApiCall()
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/check", validUserNameHandler)
	http.ListenAndServe(":8080", nil)
}
