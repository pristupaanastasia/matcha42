package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/lib/pq"
)


func LoginHandler(w http.ResponseWriter, r *http.Request){


}

func main() {

	connStr := "host=db port=5432 user=postgres password=1805 dbname=db_matcha sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/api/user/login", LoginHandler)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":9000", nil)
}