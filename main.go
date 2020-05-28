package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/danieloluwadare/golang-crud-app-db/orm"
	"github.com/gorilla/mux"
	"github.com/mailgun/mailgun-go/v3"

)

func main() {
	//initialMigration()

	// This are the method to test the jdbclike connection to the db without orm
	// to use this method below comment out the serveApplication() metthod
	// connectToMysqlDb()
	// insertIntoMysql()
	// selectFromDb()
	// querySingleRowDb()
	// allUsers()

	//serveApplication()

	//res, err:= SendSimpleMessage("sandbox.mailgun.org",
	//	"-7fba8a4e-0a6e2820")
	//if err !=nil{
	//	fmt.Println("error mail-gun", err)
	//
	//}
	//fmt.Println("result mail-gun", res)

	timeManipulation()

}

func timeManipulation(){
	todayDate := time.Now()
	fmt.Println("today date", todayDate)

	timeFormated := todayDate.Format("2006-01-02")

	fmt.Println("formatted today date", timeFormated)

	timeFormated2 := todayDate.Format("2006-01-02 15:04:05")

	fmt.Println("formatted today date", timeFormated2)

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

func SendSimpleMessage(domain, apiKey string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		"dkreativecoders@gmail.com",
		"Hello",
		"Testing some Mailgun awesomeness!",
		"danieldada123@gmail.com",
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}