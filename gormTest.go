package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

type Test struct {
	gorm.Model
	Name  string
	Email string
}

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUrl := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbName)
	fmt.Println(dbUrl)

	conn, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	initialMigration()
}

func initialMigration() {

	// Migrate the schema
	db.AutoMigrate(&User{}, &Test{})
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New User Endpoint Hit")

	var newUser User

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	// vars := mux.Vars(r)
	// name := vars["name"]
	// email := vars["email"]

	json.Unmarshal(reqBody, &newUser)

	db.Create(&newUser)
	fmt.Fprintf(w, "New User Successfully Created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Successfully Deleted User")
}
