package base

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SecretKey = "tianqi2018xbc"


func CreateToken(userId string, mac string) string {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(360)).Unix() //360个小时有效
	claims["iat"] = time.Now().Unix()
	claims["iss"] = "tq.iuoon.com"
	claims["uid"] = userId
	claims["mac"] = mac

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
