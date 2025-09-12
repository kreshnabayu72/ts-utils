package middleware

import (
	"gitea.qwertysystem.net/BETS/ts-utils/log"
	"github.com/gofiber/fiber/v2"
)

// LoggingMiddleware is a Fiber middleware that centralizes error logging.
// If you prefer calling `SendErrorLog` directly, you can skip this function.
//
// How it works:
//  1. The request enters the middleware.
//  2. The middleware calls ctx.Next() to let the route handler and other middleware run.
//  3. Inside your controller, you set ctx.Locals(ErrorLog, err) when you want
//     the error to be logged. This avoids scattering log.SendErrorLog calls everywhere.
//  4. After the handler finishes, LoggingMiddleware checks ctx.Locals(ErrorLog).
//  5. If an error is found, it sends the error message asynchronously to the remote log service.
//
// Notes:
//   - You need to pass ctxWithTimeout to usecase param as pointer.
//   - You need to set needed env to send error log.
//   - Errors are logged in a goroutine to avoid slowing down the response.
//   - If ctx.Locals(ErrorLog) is not set, nothing is logged.
//   - If you return an error directly from the handler (instead of setting ctx.Locals),
//     you may still wrap Fiber’s ErrorHandler to handle that separately.
//
// Example usage:
//
//	// 1. Register the middleware in your app initialization
//	app := fiber.New()
//	app.Use(middleware.LoggingMiddleware()) // before route initialization
//
//	// 2. In your usecase, set context value to ctxWithTimeout
//	func (u *GetUseCase) Invoke(param) error {
//	    // service calls
//
//	    if err != nil {
//	        // store the error in context for logging
//	        *param.Ctx = context.WithValue(*param.Ctx, "usecaseErr", err)
//	        return ...
//	    }
//	    return ...
//	}
//
//	// 3. In your controller, catch context value set to ctxWithTimeout
//	func GetTransaction(ctx *fiber.Ctx) error {
//	    transaction, err := usecase.FindTransactionUsecase(param)
//	    if err != nil {
//	        // catch the error from ctxWithTimeout
//	        if usecaseErr, ok := ctxWithTimeout.Value("usecaseErr").(error); ok {
//	        	ctx.Locals(util_middleware.ErrorLog, usecaseErr)
//	        }
//	        return ...
//	    }
//	    return ...
//	}
//
//	// 4. LoggingMiddleware will automatically pick up the error after ctx.Next()
//	//    and send it to the log service.
func LoggingMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		isTest := ctx.Locals(IsTest)
		if isTest != nil {
			return ctx.Next()
		}

		// Run the next handler first
		err := ctx.Next()

		// Now check error stored in locals after handler finishes
		errorLog, _ := ctx.Locals(ErrorLog).(error)
		if errorLog != nil {
			go log.SendErrorLog(errorLog.Error())
		}

		return err
	}
}
