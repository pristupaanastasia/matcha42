package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pristupaanastasia/matcha42/app/model"
	"github.com/pristupaanastasia/matcha42/app/token"
	"golang.org/x/crypto/bcrypt"

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
		var buf model.User
		buf = model.ParseJSON(w,r)
		fmt.Println(buf)
		res := model.Database.QueryRow("select * from users where login = $1 ",buf.Login)
		fmt.Println(res)
		err := res.Scan(&user.Id,&user.Email, &user.Login,
		&user.Password, &user.FirstName, &user.LastName,&user.Verify)
		if err != nil{
			fmt.Println("ошибка скан после квери")
			fmt.Println("err", err)
			fmt.Println("user", user)
			http.Redirect(w, r, model.ServerVue +"login", http.StatusUnauthorized)
			return
		}
		errno := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(buf.Password))
		fmt.Println(user.Password)
		if errno != nil{
			fmt.Println("пароль не подходит")
			http.Redirect(w, r, model.ServerVue +"login", http.StatusUnauthorized)
			return
		}
		if !isAuthenticated(w,r) {
			fmt.Println("ошибка аунтификейта")
			http.Redirect(w, r, model.ServerVue +"login", http.StatusUnauthorized)
			return
		}
		//w.WriteHeader(http.StatusOK)
		fmt.Println("redirect!! to profile")
		http.Redirect(w, r, model.Server + "/api.profile", http.StatusSeeOther)
		fmt.Printf("%s\n", model.Server +"/api.profile")
		fmt.Println("NOW REDIRECTED")
		return

	}else{
		http.ServeFile(w, r, model.ServerVue +"login")
	}
}