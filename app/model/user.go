package model

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

var Database *sql.DB
var Server string = "http://localhost:9000"

type User struct{
	Id string
	Email string
	Login string
	Password string
	FirstName string
	LastName string
	Verify bool
}

func ParseJSON(w http.ResponseWriter, r *http.Request) User{
	var user User
	decoder := json.NewDecoder(r.Body)
	decoder.UseNumber()
	err := decoder.Decode(&user)
	if err!=nil{
		log.Fatal(err)
	}
	return user
}