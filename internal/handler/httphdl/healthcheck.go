package httphdl

import (
	"net/http"

	"github.com/porrporporrpor/hydra-trust-issuer/pkg/appresponse"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (hdl *HTTPHandler) Root(ctx *fiber.Ctx) error {
	return ctx.SendString("service is running.")
}

// HealthCheck HTTPHandler
// @Tags HealthCheck
// @Method Get
// @Accept json
// Produce json
// @Success	200 {object} appresponse.responseSuccess "Success"
// @Router /healthcheck [GET]
func (hdl *HTTPHandler) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(appresponse.Success("OK"))
}

// Metrics HTTPHandler
// @Tags Metrics
// @Method Get
// @Accept json
// Produce json
// @Success	200 {string} http.StatusOK "Metrics information"
// @Router /metrics [GET]
func (hdl *HTTPHandler) Metrics() fiber.Handler {
	return adaptor.HTTPHandler(promhttp.Handler())
}
