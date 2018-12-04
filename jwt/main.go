package main

import (
	"fmt"

	"crypto/md5"
	"encoding/hex"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	key := "AllYourBase"
	ss, err := GenJwtToken(key, "test", 15000)

	fmt.Printf("%v %v", ss, err)

	fmt.Println(Md5(ss))
}

func GenJwtToken(key, body string, expire int64) (string, error) {
	mySigningKey := []byte(key)
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expire,
		Issuer:    body,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func Md5(body string) string {
	md5 := md5.New()
	md5.Write([]byte(body))
	return hex.EncodeToString(md5.Sum(nil))
}
