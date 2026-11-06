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
	commonDto "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/common/dto"
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

// sendSSERequest is the body of POST /internal/sse/send. The tenant is not in here:
// it travels in X-Client-ID, so a service token alone can never aim a push at another
// tenant's sessions.
//
// Message is the legacy free-text/name field; Event is the machine event name the newer
// consumers key on. Data carries arbitrary JSON (arrays and numbers included), which is what
// file-manager needs for {ids, purgedBy, cascadeReason, timestamp}.
type sendSSERequest struct {
	Message    string         `json:"message,omitempty"`
	Event      string         `json:"event,omitempty"`
	Reason     string         `json:"reason,omitempty"`
	Data       map[string]any `json:"data,omitempty"`
	UserIDs    []string       `json:"userIds,omitempty"`
	SessionIDs []string       `json:"sessionIds,omitempty"`
}

// SendSSE pushes one {message, reason} event to active SSE sessions in the caller's
// tenant (POST /internal/sse/send). userIDs names the recipients within that tenant;
// passing nil makes it a tenant-wide event hitting every session of that client, so
// pass one deliberately unless that is what you want.
//
// A nil receiver is a no-op, so a consumer that was constructed without a client
// (tests, or a service where notifications are optional) can call this unguarded.
//
// reason is optional and variadic only to keep it so; extra values are ignored. It
// is omitted from the body when empty.
func (c *NotificationClient) SendSSE(ctx context.Context, clientID string, userIDs []string, message string, reason ...string) error {
	if c == nil {
		return nil
	}

	payload := sendSSERequest{Message: message, UserIDs: userIDs}
	if len(reason) > 0 {
		payload.Reason = reason[0]
	}
	return c.send(ctx, clientID, payload)
}

// SendSSEWithData pushes one structured event to active SSE sessions in the caller's tenant.
// event is the machine name (e.g. "file.purged"); a Message alias with the same value is sent
// alongside it so consumers that still key on message keep working during the transition.
// data carries arbitrary JSON values, including arrays and numbers.
//
// userIDs behaves exactly as in SendSSE. A nil receiver is a no-op.
func (c *NotificationClient) SendSSEWithData(
	ctx context.Context,
	clientID string,
	userIDs []string,
	event string,
	data map[string]any,
	reason ...string,
) error {
	if c == nil {
		return nil
	}

	payload := sendSSERequest{
		Message: event,
		Event:   event,
		Data:    data,
		UserIDs: userIDs,
	}
	if len(reason) > 0 {
		payload.Reason = reason[0]
	}
	return c.send(ctx, clientID, payload)
}

// CreateNotification creates one persisted notification for a user in the caller's
// tenant through notification-service (POST /internal/). A nil receiver is a no-op.
func (c *NotificationClient) CreateNotification(ctx context.Context, clientID string, req commonDto.CreateNotificationRequest) error {
	if c == nil {
		return nil
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/internal/", bytes.NewReader(body))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set(constant.HEADER_CLIENT_ID, clientID)
	httpReq.Header.Set(constant.HEADER_SERVICE_TOKEN, c.serviceToken)

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return fmt.Errorf("notification.CreateNotification: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("notification.CreateNotification: unexpected status %d", resp.StatusCode)
	}
	return nil
}

// send marshals and posts one event body.
func (c *NotificationClient) send(ctx context.Context, clientID string, payload sendSSERequest) error {
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
