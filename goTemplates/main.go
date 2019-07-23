package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

type (
	Location struct {
		Street  string
		ZipCode string
	}

	User struct {
		Username  string
		Locations map[string]Location
	}

	UsersPage struct {
		Title string
		Users []User
	}
)

func main() {
	message, err := ioutil.ReadFile("tpl.gohtml")
	if err != nil {
		panic(err)
	}

	t, err := template.New("UsersPage").Parse(string(message))
	if err != nil {
		panic(err)
	}

	p := UsersPage{
		Title: "Users location",
		Users: []User{
			{
				Username: "johnsmith",
				Locations: map[string]Location{
					"Home": {
						Street:  "Main Street",
						ZipCode: "20192",
					},
				},
			},
		},
	}

	err = t.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
}
