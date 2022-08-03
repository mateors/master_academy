package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func main() {

	var email, role string = "billahmdmostain@gmail.com", "Admin"
	token, err := GenerateJWT(email, role)
	fmt.Println(token, err)

	ValidateJWT(token)

}

func GenerateJWT(email, role string) (string, error) {

	secretkey := "121212"
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		err = fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) {

	//jwt.ParseWithClaims()
	secretkey := "121212"
	var mySigningKey = []byte(secretkey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	fmt.Println(token, err)
	fmt.Println(token.Valid)
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println(claims, ok)

}
