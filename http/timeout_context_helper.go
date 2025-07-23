package http

import (
	"context"
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/response"
	"github.com/gofiber/fiber/v2"
)

type UseCaseFunc[T any] func(ctx context.Context) (T, *entity.HttpError)

// RunWithTimeout runs a use case with timeout and handles Fiber response to avoid rewrite in every usecase calls in controller
func RunWithTimeout[T any](
	ctx *fiber.Ctx,
	timeout time.Duration,
	useCase UseCaseFunc[T],
	successMessage string,
) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx.UserContext(), timeout)
	defer cancel()

	resultChan := make(chan T)
	errorChan := make(chan entity.HttpError)

	go func() {
		res, err := useCase(ctxWithTimeout)
		if err != nil {
			select {
			case errorChan <- *entity.ToHttpError(err):
			case <-ctxWithTimeout.Done():
			}
			return
		}
		select {
		case <-ctxWithTimeout.Done():
		case resultChan <- res:
		}
	}()

	select {
	case <-ctxWithTimeout.Done():
		return response.SendResponse(ctx, fiber.StatusRequestTimeout, nil, "Timeout")
	case err := <-errorChan:
		return response.SendResponse(ctx, err.Code, nil, err.Message)
	case res := <-resultChan:
		return response.SendResponse(ctx, fiber.StatusOK, res, successMessage)
	}
}
