package controllers

import (
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/utilities"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Interactives(c *gin.Context) {
	if err := VerifySlackRequest(c); err != nil {
		return
	}

	a, _ := io.ReadAll(c.Request.Body)

	message := models.SlackMessage{}

	err := json.Unmarshal([]byte(string(a)), &message)
	if err == nil {

	}
}

func VerifySlackRequest(c *gin.Context) error {
	timestamp := c.Request.Header["X-Slack-Request-Timestamp"][0]
	curentTime := utilities.UnixTimestampSecond(time.Now())

	timestampInt, err := strconv.ParseInt(timestamp, 10, 32)
	currentTimeInt, err := strconv.ParseInt(curentTime, 10, 32)

	if err != nil {
		return exceptions.NewBadRequestError("Timestamp is invalid")
	}

	// Check request time out
	if (currentTimeInt - timestampInt) > 60*5 {
		return exceptions.NewBadRequestError("Request time out")
	}

	requestBody, err := io.ReadAll(c.Request.Body)

	if err != nil {
		return exceptions.NewBadRequestError("Request body is invalid")
	}

	sigBaseString := "v0:" + timestamp + ":" + string(requestBody)
	slackSigningSecret := os.Getenv("SLACK_SIGNING_SECRET")

	sigKey := []byte(slackSigningSecret)
	hHash := hmac.New(sha256.New, sigKey)
	hHash.Write([]byte(sigBaseString))
	signature := "v0=" + hex.EncodeToString(hHash.Sum(nil))
	slackSignature := c.Request.Header["X-Slack-Signature"][0]

	if utilities.SecureCompare([]byte(signature), []byte(slackSignature)) {
		return exceptions.NewBadRequestError("Verify request failed!")
	}

	return nil
}
