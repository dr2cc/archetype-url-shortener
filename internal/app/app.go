package app

// 🧹🏦

import (
	"arch/config"
	handler "arch/internal/controller"
	"arch/pkg/logger"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// creating an instance of the main application
	example := App{
		InfoLogger:  logger.New("INFO: "),
		ErrorLogger: logger.New("ERROR: "),
	}

	// // 3. Define routes
	// // Раздел "ручек" и вариантов их создания

	// создаем роутер
	r := Router()

	// http.HandlerFunc— это ТИП,
	// удобный адаптер, который позволяет простой функции выполнять "контракт"
	// http.Handler на обработку HTTP-запросов (чтобы это не значило),
	// упрощая использование простых функций в качестве обработчиков.
	//
	// We wrap the plain function `greet` in http.HandlerFunc to make it a Handler
	// Мы оборачиваем простую функцию `greet` в http.HandlerFunc, чтобы сделать ее обработчиком
	gr := http.HandlerFunc(handler.Greet)

	// // HandleFunc это ФУНКЦИЯ которая регистрирует handler для заданного шаблона маршрута
	// http.HandleFunc("/", greet)// образец HandleFunc в случае использования DefaultServeMux
	// mux.HandleFunc("POST /HandleFunc", greet)
	r.HandleFunc("/HandleFunc", handler.Greet)
	//chi.HandleFunc("/HandleFunc", greet)

	// http.Handle("POST /httpHandleFunc", gr)
	// mux.Handle("POST /Handle", &example) // ❗ В таком виде работает! Что это дает пока не понял..
	r.Handle("/httpHandleFunc", gr)
	r.Handle("/Handle", &example) // ❗ В таком виде работает! Что это дает пока не понял..

	// // 4. Раздел сервера
	l.Info("The server is starting")
	//example.InfoLogger.Println("The server is starting")

	// // The handler is typically nil, in which case [DefaultServeMux] is used.
	// // Обработчик (второй параметр) по умолчанию равен nil, в этом случае используется [DefaultServeMux].
	// // Его использование не рекомендуется (можно только в простых, тестовых приложениях).
	//
	// // В рабочих приложениях следует использовать http.NewServeMux или сторонние роутеры
	// http.ListenAndServe("localhost:8080", nil)

	// Запуск сервера с обработкой ошибки
	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		l.Fatal(fmt.Errorf("app - Run - http.ListenAndServe: %w", err))
	}
}

// Тип реализующий два экземпляра логгера,
// а с методом ServeHTTP он (тип) еще и считается http.Handler
type App struct {
	InfoLogger  *logger.Logger
	ErrorLogger *logger.Logger
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.InfoLogger.Info("I use Handler!")
	fmt.Fprintln(w, "I use Handler!")
}

func Router() *chi.Mux {
	// DefaultServeMux не требует создания экземпляра роутера, только объявление его как nil (http.ListenAndServe("localhost:8080", nil))
	// mux := http.NewServeMux()
	return chi.NewRouter()
}
