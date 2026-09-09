package main

import (
	"time"
)

type Claims interface {
	SetExpireTime(ttl time.Duration)
	Valid() error
}
