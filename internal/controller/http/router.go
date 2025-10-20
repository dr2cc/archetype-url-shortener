// Package v1 implements routing paths. Each services in own file.
package http

import (
	"net/http"

	// // Swagger docs.
	// _ "arch/docs"
	"arch/internal/controller/http/middleware"
	v1 "arch/internal/controller/http/v1"
	"arch/internal/usecase"
	"arch/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(app *fiber.App, l logger.Interface, t usecase.Translation) {
	// Options
	app.Use(middleware.Logger(l))
	app.Use(middleware.Recovery(l))

	// // Prometheus - это система мониторинга, разработанная для наблюдения за распределенными системами.
	// // Он предоставляет инструменты для сбора и хранения временных рядов данных,
	// // а также для создания пользовательских запросов и алертинга на основе этих данных.
	// // Prometheus предлагает нативную поддержку для сбора метрик от приложений, что делает его идеальным выбором для мониторинга Go-приложений.
	// //
	// // Prometheus metrics
	// prometheus := fiberprometheus.New("my-service-name")
	// prometheus.RegisterAt(app, "/metrics")
	// app.Use(prometheus.Middleware)

	// // Swagger
	// app.Get("/swagger/*", swagger.HandlerDefault)

	// K8s probe
	app.Get("/healthz", func(ctx *fiber.Ctx) error { return ctx.SendStatus(http.StatusOK) })

	// Routers
	apiV1Group := app.Group("/v1")
	{
		v1.NewTranslationRoutes(apiV1Group, t, l)
	}
}
