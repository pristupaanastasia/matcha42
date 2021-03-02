package token

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/pristupaanastasia/matcha42/model"
	//"database/sql"
	"github.com/dgrijalva/jwt-go"
	"log"
	//"net/http"
	"time"
)
type Token struct{
	Id   string
	Key string
	LoginTime time.Duration
	LastSeen time.Time
}
func CreateTokenRefresh(user model.User) (string, error){
	var secretKey, error = rsa.GenerateKey(rand.Reader, 1024)
	if error != nil {
		log.Println(error)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256,jwt.MapClaims{
		"login":user.Login,
		"exp":time.Now().Add(time.Minute * 5).Unix(),
	})

	tokenString, erro:= token.SignedString(secretKey)
	if erro != nil {

		return "",erro
	}
	return tokenString, nil
}