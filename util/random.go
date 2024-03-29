package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

//RandomInt generates a random integer from min to max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//RandomString generates a string of given length
func RandomString(n int) string {
	var builder strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		builder.WriteByte(c)
	}

	return builder.String()
}

//RandomEmail returns a random email address
func RandomEmail() string {
	domains := []string{"gmail.com", "outlook.com", "icloud.com"}
	randomDomain := domains[rand.Intn(len(domains))]
	randomName := RandomString(5)
	randomEmail := fmt.Sprintf("%s@%s", randomName, randomDomain)
	return randomEmail
}

//RandomUsername returns a random 7 characters long username
func RandomUsername() string {
	return RandomString(7)
}

//RandomHashedPassword returns a hashed password and an error
func RandomHashedPassword() (string, error) {
	rawPass := RandomString(8)
	pass, err := HashPassword(rawPass)
	return pass, err
}

//RandomMovie returns a random movie name
func RandomMovie() string {
	return RandomString(12)
}
