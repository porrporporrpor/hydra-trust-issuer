package protocol

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/porrporporrpor/hydra-trust-issuer/config"
	_ "github.com/porrporporrpor/hydra-trust-issuer/docs"
	"github.com/porrporporrpor/hydra-trust-issuer/internal/handler/httphdl"
	"github.com/porrporporrpor/hydra-trust-issuer/pkg/apperror"
	"github.com/porrporporrpor/hydra-trust-issuer/pkg/middleware"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// ServeREST ...
func ServeREST() error {
	appServer := fiber.New(fiber.Config{
		DisableKeepalive: false,
		ErrorHandler:     middleware.ErrorHandler(),
	})

	appServer.Use(middleware.Recovery())
	appServer.Use(middleware.Cors())

	hdl := httphdl.NewHTTP(*config.GetConfig())

	appServer.Get("/", hdl.Root)
	appServer.Get("/metrics", hdl.Metrics())
	appServer.Get("/healthcheck", hdl.HealthCheck)
	appServer.Get("/docs/*", swagger.HandlerDefault)

	// Unprotected routes
	{
	}

	// Protected routes
	{
	}

	appServer.Use(func(ctx *fiber.Ctx) error {
		return apperror.NewPageNotFoundError()
	})

	go gracefullyShutDown(appServer)

	addr := fmt.Sprintf("%v:%v", config.GetConfig().App.Host, config.GetConfig().App.Port)
	err := appServer.Listen(addr)
	if err != nil {
		return err
	}
	return nil

}

func gracefullyShutDown(appServer *fiber.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		// Block until got a signal.
		<-c
		log.Println("Gracefully shutting down...")
		err := appServer.Shutdown()
		if err != nil {
			log.Println(fmt.Sprintf("Got error: %s while shutting down HTTP server", err.Error()))
		}
		os.Exit(0)

	}()
}
