package main

import (
	"base-api-auth/config"
	"log"
	"net/http"
)

func main() {
    // load the configuration
    cfg := config.InitConfig()

    // initialize the server with the configuration
    server := &http.Server{
        Addr:    cfg.Server.Addr,
        Handler: http.DefaultServeMux,
    }

    // start the server
    log.Printf("Server started at %s", cfg.Server.Addr)
    if err := server.ListenAndServe(); err != nil {
        log.Fatalf("Server failed: %s", err)
    }
}
