package base

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SecretKey = "tianqi2018xbc"

type CustClaims struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
	UserId    int64  `json:"uid,omitempty"`
}

func CreateToken(userId string, mac string) string {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["iss"] = "tq.iuoon.com"
	claims["uid"] = userId
	claims["mac"] = mac
	fmt.Println("duration==", time.Duration(1))
	fmt.Println("nowtime==", time.Now().Unix())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(SecretKey))
	return tokenString
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
