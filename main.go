package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// User struct hehe
type User struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("Go & MySQL!")

	db, err := sql.Open("mysql", "root:13121992@/testdb")
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	results, err := db.Query("select name from users")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.Name)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Name)
	}
}
