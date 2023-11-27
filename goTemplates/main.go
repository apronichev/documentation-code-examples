package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type Location struct {
	Street  string
	ZipCode string
}

type User struct {
	Name      string
	Locations map[string]Location
}

type UsersPage struct {
	Title string
	Users []User
}

func main() {
	message, err := os.ReadFile("templates/tpl_example.gohtml")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("UsersPage").Parse(string(message))
	if err != nil {
		panic(err)
	}

	// Define the handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Example users data
		p := UsersPage{
			Title: "Users location",
			Users: []User{
				{
					Name: "John",
					Locations: map[string]Location{
						"Home": {
							Street:  "Main Street",
							ZipCode: "20192",
						},
					},
				},
			},
		}

		// Execute the template with users data
		err := tmpl.Execute(w, p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start the server
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
