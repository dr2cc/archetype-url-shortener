# internal/app

// http.HandlerFunc— это ТИП,
	// удобный адаптер, который позволяет простой функции выполнять "контракт"
	// http.Handler на обработку HTTP-запросов (чтобы это не значило),
	// упрощая использование простых функций в качестве обработчиков.
	//
	// We wrap the plain function `greet` in http.HandlerFunc to make it a Handler
	// Мы оборачиваем простую функцию `greet` в http.HandlerFunc, чтобы сделать ее обработчиком

    // // HandleFunc это ФУНКЦИЯ которая регистрирует handler для заданного шаблона маршрута
	// http.HandleFunc("/", greet)// образец HandleFunc в случае использования DefaultServeMux
	// mux.HandleFunc("POST /HandleFunc", greet)

    //chi.HandleFunc("/HandleFunc", greet)

	// http.Handle("POST /httpHandleFunc", gr)
	// mux.Handle("POST /Handle", &example) // ❗ В таком виде работает! Что это дает пока не понял..

    //example.InfoLogger.Println("The server is starting")

	// // The handler is typically nil, in which case [DefaultServeMux] is used.
	// // Обработчик (второй параметр) по умолчанию равен nil, в этом случае используется [DefaultServeMux].
	// // Его использование не рекомендуется (можно только в простых, тестовых приложениях).
	//
	// // В рабочих приложениях следует использовать http.NewServeMux или сторонние роутеры
	// http.ListenAndServe("localhost:8080", nil)