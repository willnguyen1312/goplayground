package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	// side effect
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// User record
type User struct {
	gorm.Model
	Name  string
	Email string
}

// InitialMigration does its thing
func InitialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the DB")
	}

	defer db.Close()

	db.AutoMigrate()
}

// HelloWorld endpoint
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi world")
}

// AllUsers endpoint
func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the DB")
	}

	defer db.Close()

	var users []User

	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// AddUser endpoint
func AddUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the DB")
	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
}

// DeleteUsers endpoint
func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the DB")
	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)
}

// UpdateUser endpoint
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the DB")
	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)
	user.Email = email

	db.Save(&user)
}
