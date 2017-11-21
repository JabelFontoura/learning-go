package main

import (
	"math/rand"
	"time"
)

func random(number int) int {
	source := rand.NewSource(time.Now().UnixNano())

	return rand.New(source).Intn(number)
}
