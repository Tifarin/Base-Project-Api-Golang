package main

import (
    "base-api-auth/cmd"
    "log"
)

func main() {
    if err := cmd.RunServer(); err != nil {
        log.Fatalf("Server failed: %s", err)
    }
}
