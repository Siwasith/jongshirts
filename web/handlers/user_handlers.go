package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/oddsteam/jongshirts/internal/sessions"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/login.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, nil)
}

func AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
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
