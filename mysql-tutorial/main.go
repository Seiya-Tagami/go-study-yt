// @see https://www.youtube.com/watch?v=DWNozbk_fuk&list=PLzUGFf4GhXBL4GHXVcMMvzgtO8-WEJIoY&index=5

package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "")

	if err != nil {
		panic(err.Error())
	}

	results, err := db.Query("SELECT name FROM users")
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
	// defer db.Close()

	// insert, err := db.Query("INSERT INTO users VALUES('ELLIOT')")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()

	// fmt.Println("Successfully Connected to MySQL database")
}

// a better way to connect to mysql
// https://zenn.dev/mstn_/articles/75667657fa5aed

func connectDB() *sql.DB {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		// error handling
	}

	c := mysql.Config{
		DBName: "db",
		User: "user",
		Passwd: "password",
		Addr: "localhost:xxxx",
		Net: "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc: jst,
	}
	
	db, err := sql.Open("mysql", c.FormatDSN())

	return nil
}