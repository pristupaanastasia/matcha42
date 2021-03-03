package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"

	"github.com/pristupaanastasia/matcha42/app/model"
	"github.com/pristupaanastasia/matcha42/app/token"
	//"github.com/dgrijalva/jwt-go/request"
	"net/http"
)

func isAuthenticated(w http.ResponseWriter,r *http.Request) bool{
	cookieToken, err := r.Cookie("token")
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("token doesnt exist")
		return false
	}


	tknStr := cookieToken.Value
	claims := &token.Claims{}
	fmt.Println("tknStr",tknStr)

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(tokenbuf *jwt.Token) (interface{}, error) {
		return token.PublicKey, nil
	})
	fmt.Println("tkn", tkn)
	fmt.Println("err", err)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	if !tkn.Valid {
		fmt.Println("tkn Valid", tkn.Valid)
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true
}

func LoginHandler(next http.Handler) http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !isAuthenticated(w,r) {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func LoginUserHandler(w http.ResponseWriter,r *http.Request){
	if r.Method == "POST"{
		var user model.User
		r.ParseForm()
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")
		if !isAuthenticated(w,r) {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}
		//w.WriteHeader(http.StatusOK)
		http.Redirect(w, r,  "/", http.StatusSeeOther)



	}else{
		http.ServeFile(w, r, "auth/view/auth.html")
	}
}