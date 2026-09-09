package main

import (
	"time"
)

const (
	TOKEN_TIME   = 5 * time.Minute
	REFRESH_TIME = 1 * time.Hour
	SECRET       = "secret"
)

func main() {
	userClaims()
	personClaims()
}
