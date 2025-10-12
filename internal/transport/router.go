package router

import "github.com/go-chi/chi"

func Router() *chi.Mux {
	// DefaultServeMux не требует создания экземпляра роутера, только объявление его как nil (http.ListenAndServe("localhost:8080", nil))
	// mux := http.NewServeMux()
	return chi.NewRouter()

}
