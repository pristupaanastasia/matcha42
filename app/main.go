package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
	"log"
	"net/http"
)
var database *sql.DB

func loginHandler(next http.Handler) http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request) {

		//authToken := r.Header().Get("Authorization")
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

func verifEmail(w http.ResponseWriter, r *http.Request){
	fmt.Println("yes!",r.Context().Value("email").(string))

}

func registrHandler(next http.Handler) http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		r.ParseForm()
		login := r.FormValue("login")
		email := r.FormValue("email")
		pass := r.FormValue("password")
		first_name := r.FormValue("first_name")
		last_name := r.FormValue("last_name")
		id := uuid.New().String()


		ctx := context.WithValue(
			r.Context(), "email", email)
		ctx = context.WithValue(
			r.Context(), "login", login)
		ctx = context.WithValue(
			r.Context(), "pass", pass)
		ctx = context.WithValue(
			r.Context(), "first", first_name)
		ctx = context.WithValue(
			r.Context(), "last", last_name)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
		fmt.Fprint(w, login, pass, first_name, last_name,id,"|",len(id),"|")
		/*_, err := database.Exec("insert into users (id_user, login, email,password, first_name, last_name) values ($1, $2, $3,$4,$5,$6) ",
			id, login, email,pass,first_name,last_name)
		if err != nil {
			fmt.Println("error database",err)
			return
		}*/
	}else{
		http.ServeFile(w, r, "auth/view/registr.html")
		}
	}
	return http.HandlerFunc(fn)
}

func indexHandler(w http.ResponseWriter, r *http.Request){

	http.ServeFile(w, r, "auth/view/registr.html")
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

	database = db
	defer db.Close()

	http.Handle("/registr", recoverHandler(registrHandler(http.HandlerFunc(verifEmail))))
	http.Handle("/", loginHandler(recoverHandler(http.HandlerFunc(indexHandler))))
	fmt.Println("Server is listening...")
	http.ListenAndServe(":9000", nil)

}