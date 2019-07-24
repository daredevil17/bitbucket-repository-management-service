package main

import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Println("Go MySQL Tutorial")

    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
    log.Println(db);
    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }
    // defer the close till after the main function has finished
    // executing
    defer db.Close()

    stmtIns, err := db.Query("insert into test values('John', 'D.', 'Doe')")
	if err != nil {
		panic(err)
	}
	// Close the statement when we leave main() / the program terminates
	defer stmtIns.Close()


}
