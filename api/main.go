package main

import (
	"log"
	"net/http"

	"github.com/john-mayou/leetcli/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	log.Printf("Server running on port %s\n", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
