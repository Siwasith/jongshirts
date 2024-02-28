package handlers

import (
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/oddsteam/jongshirts/internal/sessions"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Do something with the user, maybe register them/sign them in
	_ = user
	session, _ := sessions.NewSession(r)
	// Authentication goes here
	// ...
	email := user.Email
	userImage := user.AvatarURL
	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Values["username"] = email
	session.Values["userImage"] = userImage
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}