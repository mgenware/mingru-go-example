package main

import (
	"database/sql"
	"fmt"
	"mingru-go-example/da"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open a DB connection on localhost
	db, err := sql.Open("mysql", "root:123456@/test")
	if err != nil {
		panic(err)
	}

	// Select all user profiles
	users, err := da.User.SelectAllUserProfiles(db)
	if err != nil {
		panic(err)
	}

	// Loop through the result
	for _, user := range users {
		fmt.Printf("ID: %v, Name: %v, Sig: %v\n", user.ID, user.Name, user.Sig)
	}
}
