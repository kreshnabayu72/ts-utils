package middleware_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/middleware"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
	_ "golang.org/x/image/webp"
)

func testMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(middleware.CompanyCodeKey, "companyCode")
		c.Locals(middleware.UserKey, primitive.NewObjectID())

		return c.Next()
	}
}

func TestMultiuploadMiddleware_Success_With_File(t *testing.T) {
	app := fiber.New()
	app.Use(testMiddleware())

	app.Post("/upload", middleware.MultiUploadMiddleware("test-folder", []string{"uploadImage"}), func(c *fiber.Ctx) error {
		uploadImage := c.Locals(middleware.ContextKey("uploadImage")).(string)

		assert.Equal(t, true, strings.Contains(uploadImage, "test-folder"))

		return nil
	})

	path := "../main.go"
	absPath, err := filepath.Abs(path)
	assert.Nil(t, err)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("uploadImage", absPath)
	assert.NoError(t, err)

	sample, err := os.Open(path)
	assert.NoError(t, err)

	_, err = io.Copy(part, sample)

	assert.NoError(t, writer.Close())

	req := httptest.NewRequest(fiber.MethodPost, "/upload", body)
	req.Header.Set(fiber.HeaderContentType, writer.FormDataContentType())

	result, err := app.Test(req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestMultiuploadMiddleware_Success_Without_File(t *testing.T) {
	app := fiber.New()
	app.Use(testMiddleware())

	app.Post("/upload", middleware.MultiUploadMiddleware("test-folder", []string{"uploadImage"}), func(c *fiber.Ctx) error {
		_, ok := c.Locals(middleware.ContextKey("uploadImage")).(string)
		assert.False(t, ok)

		return nil
	})

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	assert.NoError(t, writer.Close())

	req := httptest.NewRequest(fiber.MethodPost, "/upload", body)
	req.Header.Set(fiber.HeaderContentType, writer.FormDataContentType())

	result, err := app.Test(req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
