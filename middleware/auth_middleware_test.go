package middleware_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/middleware"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/helper"
)

func strPtr(s string) *string { return &s }

func newApp() *fiber.App {
	app := fiber.New()
	app.Get("/test", middleware.AuthMiddleware(), func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	return app
}

func TestAuthMiddleware(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		userID      *string
		clientID    *string
		wantStatus  int
		wantMessage string
	}{
		{
			name:       "both headers present → 200",
			userID:     strPtr("user-123"),
			clientID:   strPtr("client-456"),
			wantStatus: fiber.StatusOK,
		},
		{
			name:        "missing X-User-ID → 401",
			clientID:    strPtr("client-456"),
			wantStatus:  fiber.StatusUnauthorized,
			wantMessage: "Unauthorized",
		},
		{
			name:        "missing X-Client-ID → 401",
			userID:      strPtr("user-123"),
			wantStatus:  fiber.StatusUnauthorized,
			wantMessage: "Unauthorized",
		},
		{
			name:        "both headers missing → 401",
			wantStatus:  fiber.StatusUnauthorized,
			wantMessage: "Unauthorized",
		},
		{
			name:        "X-User-ID empty string → 401",
			userID:      strPtr(""),
			clientID:    strPtr("client-456"),
			wantStatus:  fiber.StatusUnauthorized,
			wantMessage: "Unauthorized",
		},
		{
			name:        "X-Client-ID empty string → 401",
			userID:      strPtr("user-123"),
			clientID:    strPtr(""),
			wantStatus:  fiber.StatusUnauthorized,
			wantMessage: "Unauthorized",
		},
		{
			name:        "both headers empty string → 401",
			userID:      strPtr(""),
			clientID:    strPtr(""),
			wantStatus:  fiber.StatusUnauthorized,
			wantMessage: "Unauthorized",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			if tc.userID != nil {
				req.Header.Set(constant.HEADER_USER_ID, *tc.userID)
			}
			if tc.clientID != nil {
				req.Header.Set(constant.HEADER_CLIENT_ID, *tc.clientID)
			}

			resp, err := newApp().Test(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tc.wantStatus, resp.StatusCode)

			if tc.wantStatus == fiber.StatusUnauthorized {
				// Validate Content-Type
				assert.Contains(t, resp.Header.Get("Content-Type"), fiber.MIMEApplicationJSON)

				body, err := io.ReadAll(resp.Body)
				require.NoError(t, err)

				var got helper.HttpResponse
				require.NoError(t, json.Unmarshal(body, &got),
					"response body should be valid JSON; got: %s", string(body))

				assert.Equal(t, tc.wantMessage, got.Message)
				assert.Nil(t, got.Data)
				assert.Nil(t, got.Err)
			}
		})
	}
}
