package helper

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
)

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

func GenerateSessionID() string {
	randomString, err := GenerateRandomString(16)
	if err != nil {
		fmt.Println("Error generating random string:", err)
		return ""
	}
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	return fmt.Sprintf("%s_%d", randomString, timestamp)
}
