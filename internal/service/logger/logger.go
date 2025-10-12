package logger

import (
	"log"
	"os"
)

func NewLogger(prefix string) *log.Logger {
	// скопировал у Тузова, не понимаю, что где значит
	// Вроде как os.Stdout это выходной поток (даже толком не знаю, что это)
	//return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	return log.New(os.Stdout, prefix, log.Ldate|log.Ltime)
}
