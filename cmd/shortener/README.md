В данной директории содержится код, который скомпилируется в бинарное приложение.

Рекомендуется помещать только код, необходимый для запуска приложения, но не бизнес-логику.

Название директории должно соответствовать названию приложения.

Директория `cmd/shortener` содержит:
- точку входа в приложение (функция `main`)
- инициализацию (внедрение) зависимостей (можно вынести в отдельный пакет `internal/app`)
- настройку и запуск HTTP-сервера (можно вынести в отдельный пакет `internal/router`, drk- я назвал `internal/transport`)
- обработку сигналов завершения работы приложения

Следуя лучшим практикам Go, ваш файл main.go должен содержать следующие элементы:
- Загрузка конфигурации: чтение настроек приложения, таких как порт сервера, строка подключения к базе данных и другие конфигурации из переменных среды или файла конфигурации.
- Внедрение зависимостей: инициализируйте основные компоненты вашего приложения, такие как клиент базы данных, уровень хранения и обработчики запросов, и соедините их между собой.
- Настройка маршрутизатора: создайте экземпляр маршрутизатора Chi и зарегистрируйте маршруты, передав обработчики из логического уровня вашего приложения.
- Запуск сервера: запустите HTTP-сервер, обычно с помощью http.ListenAndServe, и корректно обработайте возможные ошибки запуска.


func main() {
	// 1. Load configuration
	cfg := config.LoadConfig()

	// 2. Initialize storage (e.g., Redis or a database)
	store, err := storage.NewStore(cfg.RedisURL)
	if err != nil {
		log.Fatalf("failed to create storage: %v", err)
	}

	// 3. Initialize handlers with dependencies
	h := handlers.New(store, cfg.BaseURL)

	// 4. Set up the Chi router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Define routes
	r.Post("/shorten", h.ShortenURL)
	r.Get("/{shortURL}", h.Redirect)

	// 5. Start the server with graceful shutdown
	server := &http.Server{
		Addr:    cfg.ServerPort,
		Handler: r,
	}

	go func() {
		log.Printf("Starting server on %s", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not listen on %s: %v", cfg.ServerPort, err)
		}
	}()

	// Graceful shutdown logic
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down the server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}
	log.Println("Server stopped.")
}


your-url-shortener/
├── cmd/
│   └── url-shortener/
│       └── main.go       # The application entry point (as shown above)
├── internal/
│   ├── config/
│   │   └── config.go     # Loads and manages application settings
│   ├── handlers/
│   │   └── handlers.go   # Defines all HTTP handler functions
│   ├── storage/
│   │   └── storage.go    # Handles all data persistence logic (e.g., Redis client)
│   └── models/
│       └── url.go        # Defines data structures for request/response bodies
├── go.mod
├── go.sum
└── Dockerfile