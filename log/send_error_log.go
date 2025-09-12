package log

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func SendErrorLog(errMessage string) error {
	slackWebhook := os.Getenv("SLACK_WEBHOOK_URL")
	if slackWebhook == "" {
		log.Println("sendErrorLog: SLACK_WEBHOOK_URL not set")
		return nil
	}

	escaped := strconv.Quote(errMessage)
	raw := `{"text":` + escaped + `}`

	resp, err := http.Post(slackWebhook, "application/json", strings.NewReader(raw))
	// Don't throw error, if notif is down, let transaction continue
	if err != nil {
		fmt.Println("sendErrorLog: " + err.Error())
	}

	// Drain & close body so the connection can be reused
	_, _ = io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()

	return nil
}
