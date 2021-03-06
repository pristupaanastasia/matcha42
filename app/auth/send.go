package auth

import (
	"crypto/tls"
	"fmt"
	"github.com/pristupaanastasia/matcha42/app/model"
	"log"
	"net/smtp"
)
type Mail struct {
	senderId string
	toIds    string
	subject  string
	body     string
}


func verifyEmail(token string,email string,id string){
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
	fmt.Println("mail!!!!!" , mail.toIds)
	mail.subject = "This is the email subject"
	mail.body = " "+ model.Server +"/api.user.verify?token=" + token +  "&id=" + id + " "
	fmt.Println(mail.body)
	messageBody := mail.BuildMessage()
	fmt.Println(messageBody)

	fmt.Println("mail id: ", id)
	fmt.Println("mail t: ", token)
	smtpServer := SmtpServer{host: "smtp.gmail.com", port: "465"}
	auth := smtp.PlainAuth("", mail.senderId, "84&U4@bg5%FZ", smtpServer.host)
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