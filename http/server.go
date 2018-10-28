package main

import (
	"fmt"
	"gofullstack/basics/validationkit"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
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
	MakeApiCall()
	http.HandleFunc("/", greet)
	http.HandleFunc("/check", validUserNameHandler)
	http.ListenAndServe(":8080", nil)
}
