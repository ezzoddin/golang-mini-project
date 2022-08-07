package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"time"
)

var mySigningKey = []byte(os.Getenv("SECRET_KEY"))

func GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["client"] = "Krissanawat"
	claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	TokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong %s", err.Error())
		return "", err
	}

	return TokenString, nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJWT()
	fmt.Println(validToken)
	if err != nil {
		fmt.Println("Faiked to generate token")
	}

	fmt.Fprintf(w, string(validToken))
}

func handleRequests() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func main() {
	handleRequests()
}
