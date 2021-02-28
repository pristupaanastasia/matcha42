package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"net/smtp"
)
var database *sql.DB
var Server string = "http://localhost:9000"


type Mail struct {
	senderId string
	toIds    string
	subject  string
	body     string
}

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


func verifEmail(token string,email string,first_name string){
	// Set up authentication information.
	/*from := "gypsy_camp@mail.ru"

	// use we are sending email to
	to := email
	host := "mail.ru"
	auth := smtp.PlainAuth("", from, "password", host)
	message := token.Raw + " " + first_name
	fmt.Println(message)
	if err := smtp.SendMail(host+":25", auth, from, []string{to}, []byte(message)); err != nil {
		fmt.Println("Error SendMail: ", err)
		os.Exit(1)
	}
	fmt.Println("Email Sent!")*/
	mail := Mail{}
	mail.senderId = "anastasiapristupa1998181805@gmail.com"
	mail.toIds = email
	mail.subject = "This is the email subject"
	mail.body = " "+ Server +"/verif?token=" + token +  " "
	messageBody := mail.BuildMessage()

	smtpServer := SmtpServer{host: "smtp.gmail.com", port: "465"}
	auth := smtp.PlainAuth("", mail.senderId, "", smtpServer.host)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}
	fmt.Println("auth")
	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("dial")
	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("newclient")
	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}
	fmt.Println("Auth")
	// step 2: add all from and to
	if err = client.Mail(mail.senderId); err != nil {
		log.Panic(err)
	}
	fmt.Println("send mail")
	if err = client.Rcpt(mail.toIds); err != nil {
		log.Panic(err)
	}
	fmt.Println("data")
	// Data
	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Mail sent successfully")

}
func verifHandler(w http.ResponseWriter, r *http.Request){
	tokenString :=r.URL.Query().Get("token")
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("<YOUR VERIFICATION KEY>"), nil
	})
	fmt.Println(token)

}
func registrHandler(w http.ResponseWriter, r *http.Request){

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
		_, err := database.Exec("insert into users (id_user, login, email,password, first_name, last_name,verif) values ($1, $2, $3,$4,$5,$6, false) ",
			id, login, email,pass,firstname,lastname)
		if err != nil {
			fmt.Println("error database",err)
			return
		}
		//publicKey := &secretKey.PublicKey
		//_,er := database.Exec("insert into user_session (id_user, session_key, login_time, last_seen_time)")

	}else{
		http.ServeFile(w, r, "auth/view/registr.html")
	}

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

	database = db
	defer db.Close()

	http.Handle("/registr", recoverHandler(http.HandlerFunc(registrHandler)))
	http.Handle("/", loginHandler(recoverHandler(http.HandlerFunc(indexHandler))))
	http.Handle("/verif", loginHandler(recoverHandler(http.HandlerFunc(verifHandler))))
	fmt.Println("Server is listening...")
	http.ListenAndServe(":9000",nil)

}