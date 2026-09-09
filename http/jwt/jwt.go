package main

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateTokens(claims Claims, tokenTime, refreshTime time.Duration, secret []byte) (string, string, error) {

	claims.SetExpireTime(tokenTime)
	Token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		return "", "", errors.New("error while signing token : " + err.Error())
	}

	claims.SetExpireTime(refreshTime)
	RefreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		return "", "", errors.New("error while signing refresh token : " + err.Error())
	}
	return Token, RefreshToken, nil
}

func GetClaimsFromToken(claims Claims, token string, secret []byte) error {
	token_parsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token_parsed.Valid {
		log.Println("unauthorized : " + err.Error())
		return err
	}
	return nil
}
