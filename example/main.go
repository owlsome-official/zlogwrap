package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/owlsome-official/zlogwrap"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	// Default
	logger1 := zlogwrap.New()
	logger1.Debug()
	logger1.Debug("Debug Log")
	logger1.Info("Info Log")
	logger1.Error("Error Log")

	fmt.Println()

	// With SetField (v0.2.0+)
	myJSON := []byte(`{"num":6.13,"strs":["a","b"]}`)
	logger1.SetField("myKey", myJSON).Debug()

	// SetField x Fiber.Map (v0.2.3+)
	logger1.SetField("fiberMap", fiber.Map{
		"request_body": fiber.Map{
			"name":    "Chinnawat",
			"surname": "Chimdee",
		},
	}).Debug("")

	// a possibilities use please go to func `seeMoreUsage()`
	seeMoreUsage()

	// 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀
	// 🚀 USAGE IN FIBER APPLICATION 🚀
	// 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀 🚀

	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(requestid.New(requestid.Config{
		// zlogwrap using custom header and context key, 'Transaction-Id' instead of 'X-Request-Id'
		// Note: `zlogwrap.RequestIDHeaderKeyTag` and `zlogwrap.RequestIDContextKeyTag`
		//       are also available to use (they're constant)
		//
		// Header:     zlogwrap.RequestIDHeaderKeyTag,  // Can use this
		ContextKey: zlogwrap.RequestIDContextKeyTag, // or use ContextKey (also available too)
	}))

	app.Get("/", Handler) // GET http://localhost:8000/

	fmt.Println("Listening on http://localhost:8000")
	fmt.Println("Try to send a request :D")
	app.Listen(":8000")
}

func Handler(c *fiber.Ctx) error {

	// With Custom Logger and Context
	serviceName := "[Custom Logger]"
	myLogger := zerolog.New(os.Stdout).With().
		Str("foo", "bar").
		Float64("money", 10.99).
		Logger()
	logger5 := zlogwrap.New(zlogwrap.Config{
		ServiceName: serviceName,
		Logger:      myLogger,
		Context:     c,
	})
	logger5.Debug()
	logger5.Debug("Debug Log")
	logger5.Info("Info Log")
	logger5.Error("Error Log")

	return c.SendString("Watch your app logs!")
}

func seeMoreUsage() {
	time.Sleep(1000)
	fmt.Println()

	// You can define the service name as header of log message
	serviceName := "[Header in message]:"
	logger2 := zlogwrap.New()
	logger2.Debug()
	logger2.Debug(serviceName, "Debug Log")
	logger2.Info(serviceName, "Info Log")
	logger2.Error(serviceName, "Error Log")

	time.Sleep(1000)
	fmt.Println()

	// The better way to define the service name
	serviceNameWithConfig := "LOGGER3"
	logger3 := zlogwrap.New(zlogwrap.Config{
		ServiceName: serviceNameWithConfig,
	})
	logger3.Debug()
	logger3.Debug("Debug Log")
	logger3.Info("Info Log")
	logger3.Error("Error Log")

	time.Sleep(1000)
	fmt.Println()

	logger4 := zlogwrap.New(zlogwrap.Config{
		ServiceName: "Customize Fields",
	})
	logger4.
		SetField("bool", true).
		SetField("int", 12345).
		SetField("int64", int64(12345)).
		SetField("float64", float64(12345.01)).
		SetField("[]byte", []byte(`{"json_key": "json_value"}`)).
		SetField("[]int", []int{1, 2, 3, 4, 5}).
		SetField("[]int64", []int64{1, 2, 3, 4, 5}).
		SetField("[]float64", []float64{1.1, 2.2, 3.3, 4.4, 5.5}).
		SetField("[]string", []string{"a", "b", "c"}).
		SetField("", ""). // << expected skip this line
		SetField("str", "string").
		Debug()

	fmt.Println()

	logger5 := zlogwrap.New(zlogwrap.Config{
		ServiceName: "Customize Fields with Caller",
		Logger:      log.With().Caller().Logger(),
	})
	logger5.SetField("title", "WithCaller").GetLogEvent(zerolog.DebugLevel).Msg("a message")

	fmt.Println()
}
