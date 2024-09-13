package tests

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/owlsome-official/zlogwrap"
)

func Parsing(c *fiber.Ctx) error {
	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: "Parsing",
		Context:     c,
	})
	body := fiber.Map{}
	json.Unmarshal(c.Request().Body(), &body)

	logger = logger.SetField("before", "Just set a field and parsing to fn")
	result, _ := ReceivingLogger(logger)

	logger.SetField("body", body).SetField("result", result).Info("logging")
	return c.SendStatus(fiber.StatusNoContent)
}

func ReceivingLogger(logger zlogwrap.ZerologWrapper) (string, error) {
	mockReturn := "Something"
	logger.Info("ReceivingLogger (focus on key: before)")

	logger = logger.SetField("after", "### NO PRINTING & NO MUTATE ###")
	return mockReturn, nil
}
