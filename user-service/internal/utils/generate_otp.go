package utils

import (
	"math/rand"
	"strconv"
)


func GenerateOTP() string {
	min := 100000
	max := 999999
	return strconv.Itoa(rand.Intn(max-min) + min)
}