package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/oddsteam/jongshirts/internal/sessions"
	"github.com/oddsteam/jongshirts/web/handlers"
)

// start the web server r.HandleFunc("/books/{title}", CreateBook).Methods("POST")

func Start() {
	fmt.Println("Starting server")
	GOOGLE_CLIENT_ID := os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET := os.Getenv("GOOGLE_CLIENT_SECRET")
	CALLBACK_URL := "http://localhost:8080/auth/google/callback"
	gothic.Store = sessions.Store
	goth.UseProviders(
		google.New(GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, CALLBACK_URL, "email"),
	)
	r := mux.NewRouter()

	r.HandleFunc("/auth/{provider}", handlers.AuthHandler)
	r.HandleFunc("/auth/{provider}/callback", handlers.CallbackHandler)
	r.HandleFunc("/", handlers.HomeHandler)
	// r.HandleFunc("/cart", handlers.CartHandler).Methods("POST")
	r.HandleFunc("/showcart", handlers.ShowCart)
	r.HandleFunc("/login", handlers.LoginHandler)
	r.HandleFunc("/logout", handlers.LogoutHandler)
	r.HandleFunc("/auth", handlers.AuthenticationHandler)
	http.ListenAndServe(":8080", r)
}
