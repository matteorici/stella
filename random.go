package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func pick(values []string) string {

	return values[rand.Intn(len(values))]
}
