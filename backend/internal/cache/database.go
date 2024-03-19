package cache

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

func NewDBConnection() (*sql.DB, error) {
    
    if err := godotenv.Load(); err != nil {
        return nil, err
    }

    
    host := os.Getenv("POSTGRES_HOST")
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    dbname := os.Getenv("POSTGRES_DATABASE")
    port := os.Getenv("POSTGRES_PORT")

    
    db, err := sql.Open("postgres", "host="+host+" port="+port+" user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")
    if err != nil {
        return nil, err
    }

    return db, nil
}

func InitializeDatabase() (*sql.DB, error) {
    
    db, err := NewDBConnection()
    if err != nil {
        return nil, err
    }

    
    if err = db.Ping(); err != nil {
        db.Close()
        return nil, err
    }

    log.Println("Database connection established")
    return db, nil
}
