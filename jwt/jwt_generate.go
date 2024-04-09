package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	SECRETKEY = []byte("AllYourBase")
)

type Claims struct {
	UserID int64
	jwt.StandardClaims
}

func main() {
	tokenString := generateToken()

	// 验证 token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SECRETKEY, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}
}

func generateToken() string {
	claims := Claims{
		UserID: 12,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 64800,
			Issuer:    "may",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SECRETKEY)

	if err != nil {
		fmt.Println("生成 token 出错", err)
		return ""
	}
	fmt.Println("生成 token 成功", tokenString)
	return tokenString
}
