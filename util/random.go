package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

const littleLetters = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(littleLetters)

	for i := 0; i < n; i++ {
		sb.WriteByte(littleLetters[rand.Intn(k)])
	}
	return sb.String()
}

func RandomEmail() string {
	return RandomString(10) + "@" + RandomString(5) + ".com"
}

func RandomLongUrl() string {
	return "https://" + RandomString(10) + ".com"
}

func RandomShortUrl() string {
	return uuid.New().String()
}

func RandomUserName() string {
	return RandomString(10)
}

func RandomFullName() string {
	return RandomString(10) + " " + RandomString(10)
}

func RandomUserType() string {
	userType := []string{"FREE", "PREMIUM"}
	return userType[rand.Intn(len(userType))]
}
