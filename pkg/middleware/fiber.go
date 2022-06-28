package middleware

import (
	"net/http"
	"strconv"

	"github.com/porrporporrpor/hydra-trust-issuer/pkg/apperror"
	"github.com/porrporporrpor/hydra-trust-issuer/pkg/appresponse"

	"github.com/gofiber/fiber/v2"
	fibercors "github.com/gofiber/fiber/v2/middleware/cors"
	fiberrecovery "github.com/gofiber/fiber/v2/middleware/recover"
)

func Cors() fiber.Handler {
	return fibercors.New(fibercors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Content-Length, Accept, Authorization, X-CSRF-TOKEN",
		AllowCredentials: true,
	})
}

func Recovery() fiber.Handler {
	return fiberrecovery.New()
}

func ErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		var (
			httpStatus = http.StatusInternalServerError
			response   = appresponse.Error(
				apperror.InternalServerError,
				string(apperror.Message[apperror.InternalServerError]),
				apperror.SingleErrorMessage{Message: err.Error()},
			)
		)

		e, ok := apperror.IsAppError(err)
		if ok {
			httpStatus = FindHTTPStatusCodeFormErrorCode(e.Code)
			response = appresponse.Error(e.Code, string(e.Message), e.Errors)
		}

		return ctx.Status(httpStatus).JSON(response)
	}
}

func FindHTTPStatusCodeFormErrorCode(errCode apperror.ErrorCode) int {
	strErrCode := strconv.Itoa(int(errCode))
	rawHTTPStatus := strErrCode[:3]

	code, err := strconv.Atoi(rawHTTPStatus)
	if err != nil {
		return http.StatusInternalServerError
	}

	return code
}
