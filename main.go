package main

import (
	"log"
	"net/http"

	"github.com/danieloluwadare/golang-crud-app-db/orm"
	"github.com/gorilla/mux"
)

func main() {
	// initialMigration()

	// This are the method to test the jdbclike connection to the db without orm
	// to use this method below comment out the serveApplication() metthod
	// connectToMysqlDb()
	// insertIntoMysql()
	// selectFromDb()
	// querySingleRowDb()
	// allUsers()

	serveApplication()

}

func serveApplication() {
	router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/welcome", homeLink)
	router.HandleFunc("/user", orm.CreateNewUser).Methods("POST")

	// router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	// router.HandleFunc("/events/{id}", updateEvent).Methods("PUT")
	router.HandleFunc("/users", orm.GetAllUsers).Methods("GET")
	// router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// go mod init github.com/danieloluwadare/golang-crud-app-db
