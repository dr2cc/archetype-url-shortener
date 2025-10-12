package main

// helloweb - Snippet for sample hello world webapp (Go)
// wr		- Snippet for http Response (Go)

import (
	"arch/internal/app"
	handler "arch/internal/handlers"
	"arch/internal/service/logger"
	router "arch/internal/transport"
	"net/http"
)

// Следуя лучшим практикам Go, ваш файл main.goдолжен содержать следующие элементы:
// - Загрузка конфигурации: чтение настроек приложения, таких как порт сервера, строка подключения к базе данных и другие конфигурации из переменных среды или файла конфигурации.
// - Внедрение зависимостей: инициализируйте основные компоненты вашего приложения, такие как клиент базы данных, уровень хранения и обработчики запросов, и соедините их между собой.
// - Настройка маршрутизатора: создайте экземпляр маршрутизатора Chi и зарегистрируйте маршруты, передав обработчики из логического уровня вашего приложения.
// - Запуск сервера: запустите HTTP-сервер, обычно с помощью http.ListenAndServe, и корректно обработайте возможные ошибки запуска.

func main() {
	// // 1. Раздел экземпляра роутера

	// // 2. Раздел экземпляра "приложения" (структуры в которой определим все сущности проекта)
	// // - логгер (уже есть!). Виды экземпляров логгера вынесу в оттдельную структуру, а потом уже ее сюда
	// // - видимо и роутер сюда (http)
	// // - db
	// // - рандомайзер
	// // -

	// creating an instance of the main application
	example := app.App{
		InfoLogger:  logger.NewLogger("INFO: "),
		ErrorLogger: logger.NewLogger("ERROR: "),
	}

	// // 3. Define routes
	// // Раздел "ручек" и вариантов их создания

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
	router.Router().HandleFunc("/HandleFunc", handler.Greet)
	//chi.HandleFunc("/HandleFunc", greet)

	// http.Handle("POST /httpHandleFunc", gr)
	// mux.Handle("POST /Handle", &example) // ❗ В таком виде работает! Что это дает пока не понял..
	router.Router().Handle("/httpHandleFunc", gr)
	router.Router().Handle("/Handle", &example) // ❗ В таком виде работает! Что это дает пока не понял..

	// // 4. Раздел сервера

	example.InfoLogger.Println("The server is starting")

	// // The handler is typically nil, in which case [DefaultServeMux] is used.
	// // Обработчик (второй параметр) по умолчанию равен nil, в этом случае используется [DefaultServeMux].
	// // Его использование не рекомендуется (можно только в простых, тестовых приложениях).
	//
	// // В рабочих приложениях следует использовать http.NewServeMux или сторонние роутеры
	// http.ListenAndServe("localhost:8080", nil)

	// Запуск сервера с обработкой ошибки
	if err := http.ListenAndServe("localhost:8080", router.Router()); err != nil {
		example.ErrorLogger.Fatal(err)
	}
}
