package main

import (
    "log"
    "net/http"
    "github.com/thefinisher/thefinisher-goapp/internal/config"
    "github.com/thefinisher/thefinisher-goapp/internal/handlers"
)

func main() {
    // Load configuration
    config.LoadConfig()

    // Create a new router and attach the handler
    router := handlers.SetupRouter()

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

