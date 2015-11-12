package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO test SET NUM = ? ") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT num FROM test WHERE id = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	// Insert square numbers for 0-24 in the database
	for i := 0; i < 25; i++ {
		_, err = stmtIns.Exec(i) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	var squareNum int // we "scan" the result in here
	// Query the square-number of 13
	err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 13 is: %d\n", squareNum)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(25).Scan(&squareNum) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 1 is: %d\n", squareNum)
}
