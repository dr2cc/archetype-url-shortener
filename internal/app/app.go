package app

// 🧹🏦

import (
	handler "arch/internal/controller"
	"arch/internal/usecase/logger"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func Run() {
	// // 1. Раздел экземпляра роутера

	// // 2. Раздел экземпляра "приложения" (структуры в которой определим все сущности проекта)
	// // - логгер (уже есть!). Виды экземпляров логгера вынесу в оттдельную структуру, а потом уже ее сюда
	// // - видимо и роутер сюда (http)
	// // - db
	// // - рандомайзер
	// // -

	// creating an instance of the main application
	example := App{
		InfoLogger:  logger.NewLogger("INFO: "),
		ErrorLogger: logger.NewLogger("ERROR: "),
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

	example.InfoLogger.Println("The server is starting")

	// // The handler is typically nil, in which case [DefaultServeMux] is used.
	// // Обработчик (второй параметр) по умолчанию равен nil, в этом случае используется [DefaultServeMux].
	// // Его использование не рекомендуется (можно только в простых, тестовых приложениях).
	//
	// // В рабочих приложениях следует использовать http.NewServeMux или сторонние роутеры
	// http.ListenAndServe("localhost:8080", nil)

	// Запуск сервера с обработкой ошибки
	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		example.ErrorLogger.Fatal(err)
	}
}

// Тип реализующий два экземпляра логгера,
// а с методом ServeHTTP он (тип) еще и считается http.Handler
type App struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.InfoLogger.Println("I use Handler!")
	fmt.Fprintln(w, "I use Handler!")
}

func Router() *chi.Mux {
	// DefaultServeMux не требует создания экземпляра роутера, только объявление его как nil (http.ListenAndServe("localhost:8080", nil))
	// mux := http.NewServeMux()
	return chi.NewRouter()
}
