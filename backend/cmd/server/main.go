package main

import (
	"log"
	"net/http"

	"sandeep/lru/internal/cache"
	"sandeep/lru/pkg/server"
)

func main() {
    
    db, err := cache.InitializeDatabase()
    if err != nil {
        log.Fatalf("Error initializing database: %v", err)
    }
    defer db.Close() 

    
    cache := cache.NewLRUCache(1024, db)

    
    http.HandleFunc("/get", cache.HandleGet())
    http.HandleFunc("/set", cache.HandleSet())

    
    server := server.NewServer(":8080")

    
    log.Fatal(server.ListenAndServe())
}

