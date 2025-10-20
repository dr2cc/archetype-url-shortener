package app

// 🧹🏦

import (
	"arch/config"
	v1 "arch/internal/controller/http"
	"arch/internal/repo/persistent"
	"arch/internal/repo/webapi"
	"arch/internal/usecase/translation"
	"arch/pkg/httpserver"
	"arch/pkg/logger"
	"arch/pkg/postgres"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	translationUseCase := translation.New(
		persistent.New(pg),
		webapi.New(),
	)

	// // creating an instance of the main application
	// example := App{
	// 	InfoLogger:  logger.New("INFO: "),
	// 	ErrorLogger: logger.New("ERROR: "),
	// }

	// HTTP Server
	// handler := gin.New()
	// v1.NewRouter(handler, l, translationUseCase)
	// httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	//
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))
	v1.NewRouter(httpServer.App, l, translationUseCase)

	// Start servers
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	// // создаем роутер
	// r := Router()

	// gr := http.HandlerFunc(handler.Greet)

	// r.HandleFunc("/HandleFunc", handler.Greet)

	// r.Handle("/httpHandleFunc", gr)
	// r.Handle("/Handle", &example) // ❗ В таком виде работает! Что это дает пока не понял..

	// // // 4. Раздел сервера
	// l.Info("The server is starting")

	// // Запуск сервера с обработкой ошибки
	// if err := http.ListenAndServe("localhost:8080", r); err != nil {
	// 	l.Fatal(fmt.Errorf("app - Run - http.ListenAndServe: %w", err))
	// }
}

// // Тип реализующий два экземпляра логгера,
// // а с методом ServeHTTP он (тип) еще и считается http.Handler
// type App struct {
// 	InfoLogger  *logger.Logger
// 	ErrorLogger *logger.Logger
// }

// func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	a.InfoLogger.Info("I use Handler!")
// 	fmt.Fprintln(w, "I use Handler!")
// }

// func Router() *chi.Mux {
// 	// DefaultServeMux не требует создания экземпляра роутера, только объявление его как nil (http.ListenAndServe("localhost:8080", nil))
// 	// mux := http.NewServeMux()
// 	return chi.NewRouter()
// }
