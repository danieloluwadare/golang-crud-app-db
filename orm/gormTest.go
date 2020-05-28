package orm

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"net/http"
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

//Belongs to example
type CreditCard struct {
	gorm.Model
	UserID int
	User User
	payment  string
	paymentType string
}
//Belongs to specifying another primary key
type CreditCard2 struct {
	gorm.Model
	UserID int
	ForId int
	UID uint `gorm:"column:u_id"`
	User User `gorm:"foreignkey:UID"`
	payment  string
	paymentType string
}
var db *gorm.DB

//func init() {
//
//	e := godotenv.Load()
//	if e != nil {
//		fmt.Print(e)
//	}
//
//	username := os.Getenv("db_user")
//	password := os.Getenv("db_pass")
//	dbName := os.Getenv("db_name")
//	dbHost := os.Getenv("db_host")
//
//	dbUrl := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbName)
//	fmt.Println(dbUrl)
//
//	conn, err := gorm.Open("mysql", dbUrl)
//	if err != nil {
//		fmt.Println("Errrooorrrr ooo")
//		fmt.Print(err)
//	}
//
//	db = conn
//	initialMigration()
//}

func initialMigration() {

	// Migrate the schema
	db.AutoMigrate(&User{}, &Test{}, &CreditCard{},&CreditCard2{})
	db.Model(&CreditCard2{}).AddForeignKey("u_id", "users(id)", "RESTRICT", "RESTRICT")

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
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
