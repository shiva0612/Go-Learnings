package main

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type Person struct {
	Name   string
	UserId string
}

func (p Person) GetClaims() PersonClaims {
	pc := PersonClaims{}
	pc.UserId = p.UserId
	pc.Name = p.Name
	return pc
}

// ---------------------------- UserClaims ----------------------------
type PersonClaims struct {
	UserId string
	Name   string
	jwt.StandardClaims
}

func (pc PersonClaims) SetExpireTime(ttl time.Duration) {
	pc.ExpiresAt = time.Now().Add(ttl).Unix()
}
func (pc PersonClaims) Valid() error {
	return pc.StandardClaims.Valid()
}

// ---------------------------- userClaims ----------------------------

func personClaims() {
	person := Person{
		Name:   "John Doe",
		UserId: "123456",
	}
	//create claims - basically what is stored in JWT token
	claims := person.GetClaims()

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
