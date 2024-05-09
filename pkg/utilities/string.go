package utilities

import (
	"aio-server/pkg/constants"
	"crypto/subtle"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func IntToString(num int) string {
	return fmt.Sprintf("%d", num)
}

func BoolToString(bool bool) string {
	return fmt.Sprintf("%t", bool)
}

func SnakeCaseToHumanize(s string) string {
	words := strings.Split(s, "_")

	return strings.Join(words, " ")
}

func SecureCompare(s1, s2 string) bool {
	return subtle.ConstantTimeCompare([]byte(s1), []byte(s2)) == 1
}

func RandomToken(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)

	charset := constants.Charset
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GetAvatarUrl(key string) *string {
	var url string

	if key == "" || &key == nil {
		return nil
	}

	if os.Getenv("UPLOAD_LOCALLY_PATH") != "" {
		url = os.Getenv("UPLOAD_LOCALLY_PATH") + key
	} else {
		bucketName := os.Getenv("GCS_BUCKET_NAME")
		url = fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, key)
	}

	return &url
}
