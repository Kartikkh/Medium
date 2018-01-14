package auth

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"os"
	"time"
	"net/http"
	"strings"
)

var secret = []byte(os.Getenv("JWT_SECRET"))



type Claims struct {
	jwt.StandardClaims
	Username string
}



type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}


func NewClaims(claims jwt.StandardClaims, username string) *Claims {
	return &Claims{claims, username}
}

// Create a Auth Token
func GetToken(username string) (string){
	claims := NewClaims(jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Issuer:    "Medium",
	}, username)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(secret)

	return tokenString

}


// Validate Token

func validateToken (tokenString string) (*Claims, error){
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected Signing Method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token not valid")
	}
	return claims, nil
}


func  CheckRequest(r *http.Request) (*Claims, error) {
	auth := r.Header.Get("Authorization")

	if auth == "" {
		return nil, fmt.Errorf("authorization header is empty ")
	}

	token := strings.TrimPrefix(auth, "Token ")

	claims, err := validateToken(token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
