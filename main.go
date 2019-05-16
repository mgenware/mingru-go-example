package main

import (
	"database/sql"
	"fmt"

	"github.com/mgenware/mingru-go-example/da"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open a DB connection on localhost
	db, err := sql.Open("mysql", "root:123456@/employees?parseTime=true")
	if err != nil {
		panic(err)
	}

	// Select all user profiles
	users, err := da.Employee.SelectAll(db)
	if err != nil {
		panic(err)
	}

	// Loop through the result
	for _, user := range users {
		fmt.Printf("ID: %v, Name: %v %v, Gender: %v, Dates: %v/%v\n", user.ID, user.FirstName, user.LastName, user.Gender, user.BirthDate, user.HireDate)
	}
}
