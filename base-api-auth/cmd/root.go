package cmd

import (
	"base-api-auth/config"
	"base-api-auth/routes"
	"log"
	"net/http"
)

func RunServer() error {
    // load the configuration
    cfg := config.InitConfig()

    // get router
    router := routes.Router()

    // initialize the server with the configuration and router
    server := &http.Server{
        Addr:    cfg.Server.Addr,
        Handler: router,
    }

    // start the server
    log.Printf("Server started at %s", cfg.Server.Addr)
    return server.ListenAndServe()
}
