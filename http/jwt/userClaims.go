package main

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	Name   string
	Age    int
	Email  string
	UserId string
}

func (user User) GetClaims() UserClaims {
	uc := UserClaims{}
	uc.UserId = user.UserId
	uc.Name = user.Name
	uc.Age = user.Age
	uc.Email = user.Email
	return uc
}

// ---------------------------- UserClaims ----------------------------
type UserClaims struct {
	UserId string
	Name   string
	Age    int
	Email  string
	jwt.StandardClaims
}

func (uc UserClaims) SetExpireTime(ttl time.Duration) {
	uc.ExpiresAt = time.Now().Add(ttl).Unix()
}
func (uc UserClaims) Valid() error {
	return uc.StandardClaims.Valid()
}

// ---------------------------- userClaims ----------------------------

func userClaims() {
	user := User{
		Name:   "John Doe",
		Age:    30,
		Email:  "FwT0S@example.com",
		UserId: "123456",
	}
	//create claims - basically what is stored in JWT token
	claims := user.GetClaims()

	//generate tokens
	token, refreshToken, err := GenerateTokens(claims, TOKEN_TIME, REFRESH_TIME, []byte(SECRET))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(token)
	log.Println(refreshToken)

	//get claims
	newClaims := UserClaims{}
	GetClaimsFromToken(&newClaims, token, []byte(SECRET))
	log.Println(newClaims)
}
