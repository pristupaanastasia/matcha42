package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)
var Database *sql.DB
var Server string = "http://localhost:9000"



type SmtpServer struct {
	host string
	port string
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
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
type Token struct{
	Id   string
	Key string
	LoginTime time.Duration
	LastSeen time.Time
}

func VerifHandler(w http.ResponseWriter, r *http.Request){
	tokenString :=r.URL.Query().Get("token")
	claims := jwt.MapClaims{}
	var token Token
	id :=
	row := Database.QueryRow("select * from user_session where id_user = $1", )
	err := row.Scan(&token.Id, &token.Key, &token.LoginTime, &token.LastSeen)
	if err != nil{
		fmt.Println("error tocken")
		return
	}
	//publicKey,_ := x509.ParsePKCS1PublicKey([]byte(token.Key))
	tokenk, _ := jwt.ParseWithClaims(tokenString, claims, func(tokenk *jwt.Token) (interface{}, error) {
		return []byte(token.Key), nil
	})
	fmt.Println(token)

}
func RegistrHandler(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST" {
		r.ParseForm()
		login := r.FormValue("login")
		email := r.FormValue("email")
		pass := r.FormValue("password")
		firstname := r.FormValue("first_name")
		lastname := r.FormValue("last_name")
		id := uuid.New().String()
		var secretKey, error = rsa.GenerateKey(rand.Reader, 1024)
		if error != nil {
			log.Println(error)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodRS256,jwt.MapClaims{
			"email":email,
			"id":id,
			"exp":time.Now().Add(time.Hour).Unix(),
		})

		tokenString, erro:= token.SignedString(secretKey)
		if erro != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println(tokenString)
		verifEmail(tokenString, email,firstname)

		//ctx := context.WithValue(
		//	r.Context(), "email", email)
		//r = r.WithContext(ctx)

		/*http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
		})*/

		fmt.Fprint(w, login, pass, firstname, lastname,id,"|",len(id),"|")
		_, err := Database.Exec("insert into users (id_user, login, email,password, first_name, last_name,verif) values ($1, $2, $3,$4,$5,$6, false) ",
			id, login, email,pass,firstname,lastname)
		if err != nil {
			fmt.Println("error database",err)
			return
		}
		bytes, _ := x509.MarshalPKIXPublicKey(&secretKey.PublicKey)

		_,erl := Database.Exec("insert into user_session (id_user, session_key, login_time, last_seen_time) values($1, $2, $3,$4)", id, bytes,time.Hour ,time.Now())
		if erl != nil {
			fmt.Println("error database",err)
			return
		}
		w.WriteHeader(http.StatusOK)

	}else{
		http.ServeFile(w, r, "auth/view/registr.html")
	}

}


