package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    // Load configuration from config.json
    if err := loadConfig(); err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Create a new router
    r := mux.NewRouter()

    // Define API routes
    r.HandleFunc("/api/settings", getSettings).Methods("GET")
    r.HandleFunc("/api/settings", updateSettings).Methods("POST")
    r.HandleFunc("/api/ipfs/create", createIPFSFile).Methods("POST")
    r.HandleFunc("/api/ipfs/fetch", fetchIPFSData).Methods("POST")
    r.HandleFunc("/api/clickhouse/data", getClickHouseData).Methods("GET")

    // Serve frontend files from the "frontend/dist" directory
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist")))

    // Start the server on port 8080
    log.Println("Backend server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
