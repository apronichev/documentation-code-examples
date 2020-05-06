package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strings"
)

type User struct {
	ID         int
	firstName  string
	middleName string
	lastName   string
}

const (
	dsn       = "postgres://guest:guest@localhost:54333/guest?sslmode=disable"
	userQuery2 = "SELECT first_name FROM actor WHERE actor_id = 1;" +
		"SELECT last_name FROM actor WHERE actor_id = 1;"
)

func main() {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var firstName []string
	err = db.Select(&firstName, userQuery2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("First name: %#v\n", strings.Join(firstName, ""))
}
