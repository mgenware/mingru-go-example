package main

import (
	"database/sql"
	"fmt"
	"mingru-go-example/da"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/test")
	if err != nil {
		panic(err)
	}

	users, err := da.User.SelectAllUserProfiles(db)
	if err != nil {
		panic(err)
	}
	for _, u := range users {
		fmt.Printf("ID: %v, Name: %v, Sig: %v\n", u.UserID, u.UserName, u.UserSig)
	}
}
