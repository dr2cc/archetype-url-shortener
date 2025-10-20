package main

// helloweb - Snippet for sample hello world webapp (Go)
// wr		- Snippet for http Response (Go)

import (
	"arch/config"
	"arch/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
