package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/oddsteam/jongshirts/internal/db"
	"github.com/oddsteam/jongshirts/internal/sessions"
	"golang.org/x/vuln/client"
)

type ShirtPageData struct {
	PageTitle string
	ShirtList []ShirtList
	Username  string
	OrderTotal string
}

type ShirtList struct {
	Id    int
	Name  string
	Size  string
	Price string
	Color string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	client := db.NewClient()
	// session, err := store.Get(r, "sessions")
	session, err := sessions.NewSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if session.Values["username"] == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	tmpl, err := template.ParseFiles("web/templates/home.html")
	if err != nil {
		fmt.Println(err)
	}
	username := session.Values["username"].(string)
	var cursor uint64=0 
	orderTotal := "0"
	keys , _, _ := client.Scan(cursor,"*",100).Result()

	for key := range keys{
		client.LLen(key)
	}
	fmt.Println()


	data := ShirtPageData{
		PageTitle: "Home page",
		ShirtList: []ShirtList{
			{Id: 1, Name: "shirt 1", Price: "100", Color: "Red", Size: "XL"},
			{Id: 2, Name: "shirt 2", Price: "50", Color: "Green", Size: "L"},
			{Id: 3, Name: "shirt 3", Price: "300", Color: "Blue-green", Size: "S"},
			{Id: 4, Name: "shirt 4", Price: "77", Color: "Black", Size: "XXXL"},
		},
		Username: username,
		OrderTotal: orderTotal,
	}

	tmpl.Execute(w, data)

}