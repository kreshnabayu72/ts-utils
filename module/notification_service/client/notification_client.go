// Package client holds outbound calls to notification-service, so every producer of a
// notification or an SSE event shares one implementation instead of copying it.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
)

// NotificationClient pushes events through notification-service, which owns the live SSE
// sessions.
type NotificationClient struct {
	serviceToken string
	baseURL      string
	http         *http.Client
}

func NewNotificationClient(baseURL, serviceToken string) *NotificationClient {
	return &NotificationClient{
		serviceToken: serviceToken,
		baseURL:      baseURL,
		// A request usually blocks on this call, so it gets a deadline rather than hanging.
		http: &http.Client{Timeout: 5 * time.Second},
	}
}

// SendSSE pushes one {message, reason} event to a target user's active SSE sessions
// (POST /internal/sse/send). The target user is named by X-User-ID, per that route's
// contract — it is the push's target, not the caller's own identity.
//
// A nil receiver is a no-op, so a consumer that was constructed without a client
// (tests, or a service where notifications are optional) can call this unguarded.
//
// reason is optional and variadic only to keep it so; extra values are ignored. It
// is omitted from the body when empty.
func (c *NotificationClient) SendSSE(ctx context.Context, clientID, targetUserID, message string, reason ...string) error {
	if c == nil {
		return nil
	}

	payload := map[string]string{"message": message}
	if len(reason) > 0 && reason[0] != "" {
		payload["reason"] = reason[0]
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/internal/sse/send", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(constant.HEADER_CLIENT_ID, clientID)
	req.Header.Set(constant.HEADER_USER_ID, targetUserID)
	req.Header.Set(constant.HEADER_SERVICE_TOKEN, c.serviceToken)

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("notification.SendSSE: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("notification.SendSSE: unexpected status %d", resp.StatusCode)
	}
	return nil
}
