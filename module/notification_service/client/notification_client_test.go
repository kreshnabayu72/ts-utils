package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendSSEPostsLegacyMessageBody(t *testing.T) {
	var (
		gotBody   map[string]any
		gotClient string
		gotToken  string
	)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotClient = r.Header.Get("X-Client-ID")
		gotToken = r.Header.Get("X-Service-Token")
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := NewNotificationClient(srv.URL, "secret")
	require.NoError(t, client.SendSSE(context.Background(), "acme", []string{"7"}, "user.updated", "duplicate_login"))

	require.Equal(t, "acme", gotClient)
	require.Equal(t, "secret", gotToken)
	require.Equal(t, "user.updated", gotBody["message"])
	require.Equal(t, "duplicate_login", gotBody["reason"])
	require.Nil(t, gotBody["data"])
	require.Nil(t, gotBody["event"])
}

func TestSendSSEWithDataPostsStructuredPayload(t *testing.T) {
	var gotBody map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := NewNotificationClient(srv.URL, "secret")
	require.NoError(t, client.SendSSEWithData(
		context.Background(),
		"acme",
		nil,
		"file.purged",
		map[string]any{
			"ids":           []string{"2077960651539484700"},
			"purgedBy":      "System",
			"cascadeReason": "retention_expired",
			"timestamp":     int64(1790380799000),
		},
		"retention_expired",
	))

	require.Equal(t, "file.purged", gotBody["event"])
	require.Equal(t, "file.purged", gotBody["message"], "message alias keeps old consumers working")
	require.Equal(t, "retention_expired", gotBody["reason"])

	data, ok := gotBody["data"].(map[string]any)
	require.True(t, ok, "data must be a JSON object: %#v", gotBody["data"])
	require.Equal(t, "retention_expired", data["cascadeReason"])
	require.Equal(t, "System", data["purgedBy"])
	require.Equal(t, float64(1790380799000), data["timestamp"])
	ids, ok := data["ids"].([]any)
	require.True(t, ok)
	require.Equal(t, "2077960651539484700", ids[0])
}
