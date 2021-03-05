package main

import (
	"crypto/rand"
	"crypto/rsa"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"

	"github.com/pristupaanastasia/matcha42/app/auth"
	"github.com/pristupaanastasia/matcha42/app/model"
	"github.com/pristupaanastasia/matcha42/app/token"
)

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}




func indexHandler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w,"index")
}
func main() {
	token.PrivateKey, _ = rsa.GenerateKey(rand.Reader, 1024)
	token.PublicKey = &(token.PrivateKey.PublicKey)
	connStr := "host=db port=5432 user=postgres password=1805 dbname=db_matcha sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	model.Database = db
	defer db.Close()


	http.Handle("/register", recoverHandler(http.HandlerFunc(auth.RegistrationHandler)))
	http.Handle("/login", recoverHandler(http.HandlerFunc(auth.LoginUserHandler)))
	http.Handle("/verify", recoverHandler(http.HandlerFunc(auth.VerifyHandler)))
	http.Handle("/", auth.LoginHandler(recoverHandler(http.HandlerFunc(indexHandler))))
	fmt.Println("Server is listening...")
	http.ListenAndServe(":9000",nil)

}