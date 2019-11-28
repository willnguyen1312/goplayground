package main

import (
	"fmt"
	"log"
	"net/http"
	"playground/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Go ORM Tutorial")

	handlers.InitialMigration()
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", handlers.HelloWorld).Methods("GET")
	myRouter.HandleFunc("/users", handlers.AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", handlers.AddUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", handlers.AddUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", handlers.AddUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
