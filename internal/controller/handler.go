package handler

import (
	"fmt"
	"net/http"
	"time"
)

// Это и есть "простая" функция (the plain function) в качестве обработчика
func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
