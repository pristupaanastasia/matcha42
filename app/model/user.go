package model

import "database/sql"

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