package main

import (
	"fmt"
	"gin-api-template/internal/data"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CompareHash(password string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err==nil{
		return true, nil
	}else if err.Error()=="crypto/bcrypt: hashedPassword is not the hash of the given password"{
		return false, nil
	}
	return false, err
}

func GenerateJWT(jwtkey []byte, user *data.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["organization"] = user.Organization
	claims["email"] = user.Email

	tokenString, err := token.SignedString(jwtkey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func extractClaims(jwtkeybytes []byte,tokenStr string) (jwt.MapClaims, bool) {
	hmacSecret := jwtkeybytes
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}


