package middleware

import (
	"gitea.qwertysystem.net/BETS/ts-utils/log"
	"github.com/gofiber/fiber/v2"
)

// LoggingMiddleware is a Fiber middleware that sends error logs to a remote logging service.
// It checks if the request is a test request using the "isTest" context key; if so, it
// skips logging. Otherwise, it retrieves an error from the "errorLog" context key and
// sends it asynchronously to a logging service if an error exists.
//
// The function uses a goroutine to avoid blocking the main request-response cycle,
// ensuring a fast response time.
func LoggingMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		isTest := ctx.Locals(IsTest)
		if isTest != nil {
			return ctx.Next()
		}

		errorLog := ctx.Locals(ErrorLog).(error)

		go func(e error) {
			if e != nil {
				log.SendErrorLog(errorLog.Error())
			}
		}(errorLog)

		return ctx.Next()
	}
}
