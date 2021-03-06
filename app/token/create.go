package token

import (
	"github.com/pristupaanastasia/matcha42/app/model"
	//"crypto/rand"
	"crypto/rsa"
	//"database/sql"
	"github.com/dgrijalva/jwt-go"
	//"log"
	//"net/http"
	"time"
)
type Token struct{
	Id   string
	Key string
	LoginTime time.Time
	LastSeen time.Time
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
var PrivateKey    *rsa.PrivateKey
var PublicKey    *rsa.PublicKey
func CreateTokenRefresh(user model.User) (string, error){



	token := jwt.NewWithClaims(jwt.SigningMethodRS256,jwt.MapClaims{
		"login":user.Login,
		"exp":time.Now().Add(time.Minute * 5).Unix(),
	})

	tokenString, erro:= token.SignedString(PrivateKey)
	if erro != nil {

		return "",erro
	}
	return tokenString, nil
}