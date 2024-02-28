package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/oddsteam/jongshirts/internal/db"
	"github.com/oddsteam/jongshirts/internal/sessions"
)

type CartList struct {
	NameShirt string
	Count int
}



func CartHandler(w http.ResponseWriter, r *http.Request) {
	client := db.NewClient()
	session, _ := sessions.NewSession(r)
	username := session.Values["username"]
	// var SelectedShirts []string

	r.ParseForm()
	// ctx := context.Background()

	for key, _ := range r.Form {
		client.LPush(username.(string), key)
	}
	http.Redirect(w, r, "/showcart", http.StatusSeeOther)
}

func ShowCart(w http.ResponseWriter, r *http.Request) {
	client := db.NewClient()
	tmpl, err := template.ParseFiles("web/templates/cart.html")
	if err != nil {
		fmt.Println(err)
	}

	session, _ := sessions.NewSession(r)
	username := session.Values["username"]
	// nameShirt := client.LRange(name).
	

	data, err := client.LRange(username.(string), 0, -1).Result()
	if err != nil {
		fmt.Println(err)
	}
	
	k := make(map[string]int)

	for _, u := range data{
	k[u]+=1
	}

	tmpl.Execute(w, k)
	
}
