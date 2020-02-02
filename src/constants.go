package main

//SlackChatPostMessageURL is the full URL for the chat.postMessage API
const SlackChatPostMessageURL string = "https://slack.com/api/chat.postMessage"

//SlackTimestampHeader is the header key holding the request timestamp from Slack
const SlackTimestampHeader string = "X-Slack-Request-Timestamp"

//SlackSignatureHeader is the header key holding the HMAC signature from Slack
const SlackSignatureHeader string = "X-Slack-Signature"
