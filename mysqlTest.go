package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func connectToMysqlDb() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testgodb")

	err1 := db.Ping()

	if err1 != nil {
		log.Fatal(err1)
		panic(err.Error())

		// do something here
	}
	// if there is an error opening the connection, handle it

	fmt.Println("No Error Occured")

	// defer the close till after the main function has finished
	// executing

	defer db.Close()

}

func insertIntoMysql() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testgodb")

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Print("Error Occured")
		panic(err.Error())
	}

	fmt.Println("No Error Occured")

	// defer the close till after the main function has finished
	// executing

	defer db.Close()

	insert, err := db.Query("INSERT INTO mytable (name) VALUES ('TEST')")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}

func selectFromDb() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testgodb")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM mytable")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	tags := make([]Tag, 0)
	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.Name)
		tags = append(tags, tag)
	}

	fmt.Println(tags)

}

func querySingleRowDb() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testgodb")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	var tag Tag
	// Execute the query
	err = db.QueryRow("SELECT id, name FROM mytable where id = ?", 2).Scan(&tag.ID, &tag.Name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println(tag)

}
