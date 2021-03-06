package auth

import (
	"github.com/pristupaanastasia/matcha42/app/model"
	//"crypto/rand"
	//"crypto/rsa"
	"strings"

	"github.com/pristupaanastasia/matcha42/app/token"
	//"crypto/x509"
	_ "database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	//"log"
	"net/http"
	"time"
)


type SmtpServer struct {
	host string
	port string
}


type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}
func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderId)
	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body

	return message
}


func VerifyHandler(w http.ResponseWriter, r *http.Request){
	var tokenGet token.Token
	var user model.User

	tokenString :=r.URL.Query().Get("token")
	id :=r.URL.Query().Get("id")
	fmt.Println(id)
	fmt.Println(tokenString)
	rowUser := model.Database.QueryRow("select * from users where id_user = $1",id )
	erro := rowUser.Scan(&user.Id, &user.Email, &user.Login, &user.Password,
		&user.FirstName,&user.LastName,&user.Verify)
	if erro != nil{
		fmt.Println("error tocken user",erro)
		return
	}
	if &user == nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	row := model.Database.QueryRow("select * from user_session where id_user = $1",id )
	err := row.Scan(&tokenGet.Id, &tokenGet.Key, &tokenGet.LoginTime, &tokenGet.LastSeen)
	if err != nil{
		fmt.Println("error tocken tocken", err)
		return
	}
	if strings.TrimSpace(tokenGet.Key) != strings.TrimSpace(tokenString){
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("tokenGet.Key != tokenString")
		fmt.Println(tokenString+"|")
		fmt.Println(tokenGet.Key+"|")
		return
	}
	_,error := model.Database.Exec("update users set verify = true where id_user = $1",user.Id)
	if error != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("model.Database.Exec",error)
		return
	}
	tokenRefresh, erno := token.CreateTokenRefresh(user)
	if erno != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("token.CreateTokenRefresh",erno)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenRefresh,
		Path:    "/",
	})

	w.WriteHeader(http.StatusOK)
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}
func RegistrationHandler(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST" {
		var user model.User
		r.ParseForm()
		user.Login = r.FormValue("login")
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")
		user.FirstName = r.FormValue("first_name")
		user.LastName = r.FormValue("last_name")
		user.Id = uuid.New().String()


		tokenS := jwt.NewWithClaims(jwt.SigningMethodRS256,jwt.MapClaims{
			"email":user.Email,
			"id":user.Id,
			"exp":time.Now().Add(time.Hour).Unix(),
		})

		tokenString, erro:= tokenS.SignedString(token.PrivateKey)
		if erro != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println(tokenString)
		verifyEmail(tokenString, user.Email,user.Id)

		//ctx := context.WithValue(
		//	r.Context(), "email", email)
		//r = r.WithContext(ctx)

		/*http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
		})*/

		fmt.Fprint(w, user.Login, user.Password, user.FirstName, user.LastName,user.Id,"|",len(user.Id),"|")
		_, err := model.Database.Exec("insert into users (id_user, login, email,password, first_name, last_name,verify) values ($1, $2, $3,$4,$5,$6, false) ",
			user.Id, user.Login, user.Email,user.Password,user.FirstName,user.LastName)
		if err != nil {
			fmt.Println("error database user",err)
			return
		}
	//	bytes, _ := x509.MarshalPKIXPublicKey(&secretKey.PublicKey)

		_,erl := model.Database.Exec("insert into user_session (id_user, session_key, login_time, last_seen_time) values($1, $2, $3,$4)", user.Id, tokenString,time.Now().Add(time.Minute * 30)  ,time.Now())
		if erl != nil {
			fmt.Println("error database token",err)
			return
		}
		w.WriteHeader(http.StatusOK)
		//http.Redirect(w, r, model.Server + "/", http.StatusSeeOther)
		//тут еще будет страничка что письмо пришло на почту


	}else{
		http.ServeFile(w, r, "auth/view/register.html")
	}

}


