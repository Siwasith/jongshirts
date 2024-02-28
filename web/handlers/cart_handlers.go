package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/oddsteam/jongshirts/internal/db"
)

func CartHandler(w http.ResponseWriter, r *http.Request) {
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
