package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"strings"
)

func getSigningSecret() []byte {
	return []byte(GetEnv("SLACK_SIGNING_SECRET", ""))
}

//VerifySigningSignature verifies a signature from Slack
func VerifySigningSignature(timestamp string, requestBody []byte, refSignature []byte) bool {
	signingSecret := getSigningSecret()
	sigBaseString := []byte(strings.Join([]string{"v0", timestamp, string(requestBody)}, ":"))
	macBuffer := hmac.New(sha256.New, signingSecret)
	macBuffer.Write(sigBaseString)
	computedSignature := append([]byte("v0="), macBuffer.Sum(nil)...)
	return hmac.Equal(computedSignature, refSignature)
}
