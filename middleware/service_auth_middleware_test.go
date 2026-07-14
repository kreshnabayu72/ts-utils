package middleware

import (
	"net/http/httptest"
	"testing"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"

	"github.com/gofiber/fiber/v2"
)

func runServiceAuth(t *testing.T, secret, presented string) (int, bool) {
	t.Helper()

	reached := false
	app := fiber.New()
	app.Get("/internal/x", RequireServiceToken(secret), func(c *fiber.Ctx) error {
		reached = true
		return c.SendStatus(fiber.StatusOK)
	})

	req := httptest.NewRequest("GET", "/internal/x", nil)
	if presented != "" {
		req.Header.Set(constant.HEADER_SERVICE_TOKEN, presented)
	}
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	return resp.StatusCode, reached
}

func TestRequireServiceToken(t *testing.T) {
	if status, reached := runServiceAuth(t, "s3cret", "s3cret"); status != fiber.StatusOK || !reached {
		t.Errorf("correct token should pass: status=%d reached=%v", status, reached)
	}

	if status, reached := runServiceAuth(t, "s3cret", "wrong"); status != fiber.StatusUnauthorized || reached {
		t.Errorf("wrong token must be refused: status=%d reached=%v", status, reached)
	}

	if status, reached := runServiceAuth(t, "s3cret", ""); status != fiber.StatusUnauthorized || reached {
		t.Errorf("missing token must be refused: status=%d reached=%v", status, reached)
	}
}

// An unset secret must REFUSE, not wave everything through. A deployment that forgot to configure
// the token has to break loudly — silently unlocking the door that guards every user's permissions
// is the worst possible failure mode.
func TestRequireServiceToken_UnsetSecretFailsClosed(t *testing.T) {
	status, reached := runServiceAuth(t, "", "anything")
	if reached {
		t.Error("an unconfigured service token must not let internal requests through")
	}
	if status == fiber.StatusOK {
		t.Errorf("status = %d, want a failure status when no secret is configured", status)
	}
}
