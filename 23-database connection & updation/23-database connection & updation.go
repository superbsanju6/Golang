package main

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main()  {

	fmt.Println("Go Connect to mysql -testdb database")
	db, err :=sql.Open("mysql","root:@(127.0.0.1)/testdb")

	if err!=nil {
		panic(err.Error())
	}
	fmt.Println("Connected")

	defer db.Close()



	insert, err := db.Query("INSERT  INTO users VALUES ('ELLi')")

	if err!=nil {
		panic(err.Error())
	}

	defer  insert.Close()

	fmt.Println("Successfully inserted into user tables")

}
