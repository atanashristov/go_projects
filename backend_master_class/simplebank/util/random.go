package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	// Global rand this is deprecated with go 1.20:
	// rand.Seed(time.Now().UnixNano())
}

// Generate random integer
func RandomInt(min, max int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Int63n(max-min+1)
}

// Generate random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generate random currency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return currencies[r.Intn(n)]
}

// Generate random money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Generate random owner
func RandomOwner() string {
	return RandomString(6)
}
