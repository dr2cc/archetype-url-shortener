package app

// ❗Здесь внедрениие зависимостей

import (
	"fmt"
	"log"
	"net/http"
)

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
