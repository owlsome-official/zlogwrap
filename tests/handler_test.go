package tests

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/owlsome-official/zlogwrap"
	"github.com/stretchr/testify/assert"
)

func TestHandlerContext(t *testing.T) {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Set(zlogwrap.RequestIDHeaderKeyTag, "TestHandlerContext")
		return c.Next()
	})
	app.Get("/", HandlerContext)
	req := httptest.NewRequest("GET", "/", strings.NewReader("{\"key\": \"value\"}"))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusNoContent, resp.StatusCode, "Wrong Status Code")
}
