package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

var response = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	// TODO - load templates here
	// There are 5 templates: welcom.html, form.html, list.html, thanks.html, sorry.html
	templateNames := [5]string{"welcome", "form", "list", "thanks", "sorry"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

// welcomHandler handles root URL

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	templates["welcom"].Execute(w, nil)
}

// ListHandler handles "/list" URL
func listHandler(w http.ResponseWriter, r *http.Request) {
	templates["list"].Execute(w, response)
}

// type formData struct {
// 	*Rsvp
// 	Errors []string
// }

// FormHandler handles "/form" URL
func formHandler(w http.ResponseWriter, r *http.Request) {
	// TODO

}

func main() {
	loadTemplates()

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
