package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/mgenware/mingru-go-example/da"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open a DB connection on localhost
	db, err := sql.Open("mysql", "root:123456@/employees?parseTime=true")
	if err != nil {
		panic(err)
	}

	// Select some paged users
	users, hasNext, err := da.Employee.SelectPaged(db, 2, 10)
	if err != nil {
		panic(err)
	}

	// Loop through the result
	log.Printf("hasNext: %v\n", hasNext)
	for _, user := range users {
		fmt.Printf("ID: %v, Name: %v %v, Gender: %v, Dates: %v/%v\n", user.ID, user.FirstName, user.LastName, user.Gender, user.BirthDate, user.HireDate)
	}
}
