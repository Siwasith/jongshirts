package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oddsteam/jongshirts/internal/db"
	"github.com/oddsteam/jongshirts/internal/sessions"
	"github.com/oddsteam/jongshirts/web/handlers"
)

type Incart struct {
	Name []string
}

// start the web server r.HandleFunc("/books/{title}", CreateBook).Methods("POST")

func Start() {
	fmt.Println("Starting server")
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/cart", cartHandler).Methods("POST")
	r.HandleFunc("/showcart", ShowCart)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/auth", authenticationHandler)
	http.ListenAndServe(":8080", r)
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
	client := db.NewClient()
	// var SelectedShirts []string

	r.ParseForm()
	// ctx := context.Background()

	for key, _ := range r.Form {
		client.LPush("selectedShirt", key)
	}

	http.Redirect(w, r, "/showcart", http.StatusSeeOther)
}

func ShowCart(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/detail.html")
	if err != nil {
		fmt.Println(err)
	}

	client := db.NewClient()
	data, err := client.LRange("selectedShirt", 0, -1).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)

	tmpl.Execute(w, data)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/login.html")
	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, nil)

}

func authenticationHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.NewSession(r)
	// Authentication goes here
	// ...
	email := r.FormValue("email")
	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Values["username"] = email
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
