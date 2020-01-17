package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func getSigningSecret() []byte {
	return []byte(GetEnv("SLACK_SIGNING_SECRET", ""))
}

//VerifySigningSignature verifies a signature from Slack
func VerifySigningSignature(timestamp string, requestBody []byte, refSignature []byte) bool {
	signingSecret := getSigningSecret()
	sigBaseString := fmt.Sprintf("v0:%v:%v", timestamp, string(requestBody))
	macBuffer := hmac.New(sha256.New, signingSecret)
	macBuffer.Write([]byte(sigBaseString))
	computedSignature := append([]byte("v0="), hex.EncodeToString(macBuffer.Sum(nil))...)
	return hmac.Equal(computedSignature, refSignature)
}
