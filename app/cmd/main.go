package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"

	"github.com/pristupaanastasia/matcha42/auth"
	"github.com/pristupaanastasia/matcha42/model"
)





func loginHandler(next http.Handler) http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request) {
		/*if !isAuthenticated(r) {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}*/
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
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

	//http.ServeFile(w, r, "auth/view/registr.html")
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

	model.Database = db
	defer db.Close()

	http.Handle("/registr", recoverHandler(http.HandlerFunc(auth.RegistrationHandler)))
	http.Handle("/", loginHandler(recoverHandler(http.HandlerFunc(indexHandler))))
	http.Handle("/verif", loginHandler(recoverHandler(http.HandlerFunc(auth.VerifyHandler))))
	fmt.Println("Server is listening...")
	http.ListenAndServe(":9000",nil)

}