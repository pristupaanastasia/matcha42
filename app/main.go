package main

import (
	"crypto/rand"
	"crypto/rsa"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/pristupaanastasia/matcha42/app/auth"
	"github.com/pristupaanastasia/matcha42/app/model"
	"github.com/pristupaanastasia/matcha42/app/profile"
	"github.com/pristupaanastasia/matcha42/app/token"
	"github.com/rs/cors"
	"log"
	"net/http"
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

func main() {

	token.PrivateKey, _ = rsa.GenerateKey(rand.Reader, 1024)
	token.PublicKey = &(token.PrivateKey.PublicKey)
	token.PrivateKeyR, _ = rsa.GenerateKey(rand.Reader, 1024)
	token.PublicKeyR = &(token.PrivateKeyR.PublicKey)
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
	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
		AllowCredentials: true,
		AllowedMethods: []string{
			http.MethodGet,//http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{
			"*",//or you can your header key values which you are using in your application

		},
	})

	router.Handle("/api.user.register", recoverHandler(http.HandlerFunc(auth.RegistrationHandler)))
	router.Handle("/api.user.login", recoverHandler(http.HandlerFunc(auth.LoginUserHandler)))
	router.Handle("/api.user.verify", recoverHandler(http.HandlerFunc(auth.VerifyHandler)))
	router.Handle("/api.profile", auth.LoginHandler(recoverHandler(http.HandlerFunc(profile.IndexHandler))))
	//router.HandleFunc("api/user/auth",auth.isAuthenticated)
	handler := c.Handler(router)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":9000",handler)

}