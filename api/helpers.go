package mobileappreactnative

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateRandomString generates a random string of the specified length
func GenerateRandomString(length int) (string, error) {
	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
	var bytes []byte
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		bytes = append(bytes, digits[n.Int64()])
	}
	return string(bytes), nil
}

// GenerateSHA256 generates a SHA256 hash of the specified string
func GenerateSHA256(input string) (string, error) {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:]), nil
}

// GenerateJWT generates a JWT token based on the provided credentials
func GenerateJWT(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// VerifyJWT verifies the provided JWT token
func VerifyJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

// GetIP gets the IP address from the incoming request
func GetIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return strings.TrimSpace(ip)
}

// GetUserID generates a unique user ID based on the provided email and a timestamp
func GetUserID(email string) string {
	return fmt.Sprintf("user:%s:%d", email, time.Now().Unix())
}

// IsDebug checks if the application is running in debug mode
func IsDebug() bool {
	return os.Getenv("DEBUG_MODE") == "true"
}

// GetNextPage generates the next page URL based on the provided URL and page number
func GetNextPage(url string, page int) string {
	return fmt.Sprintf("%s?page=%d", url, page+1)
}

// GetPreviousPage generates the previous page URL based on the provided URL and page number
func GetPreviousPage(url string, page int) string {
	return fmt.Sprintf("%s?page=%d", url, page-1)
}

// GetRandomInt generates a random integer within the specified range
func GetRandomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

// GetFormattedTime formats the provided time into a human-readable string
func GetFormattedTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// GetEnvString gets the environment variable with the specified key, defaulting to the provided value
func GetEnvString(key string, defaultValue string) string {
	return os.Getenv(key) != "" ? os.Getenv(key) : defaultValue
}

// GetEnvInt gets the environment variable with the specified key, defaulting to the provided value
func GetEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Error parsing environment variable %s: %v", key, err)
		return defaultValue
	}
	return i
}