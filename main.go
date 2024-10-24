package main

import (
    "log"
    "net/http"
)

func main() {
    // Serve files in the current directory
    fs := http.FileServer(http.Dir("./hls/"))
    
    // Handler for the /hls route with CORS headers
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Set CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        // Handle preflight requests
        if r.Method == http.MethodOptions {
            return // End the request for preflight
        }

        // Log the requested URL for debugging
        log.Println("Serving request:", r.URL.Path)

        // Serve the requested file
        fs.ServeHTTP(w, r)
    })

    // Start the server
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
