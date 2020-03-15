package main

import (
    "log"
    "fmt"

    "github.com/joho/godotenv"

    _ "github.com/d1rtyvans/blubird/app"    // TODO: May not need
    _ "github.com/d1rtyvans/blubird/config" // TODO: Set this up with env variables
    "github.com/d1rtyvans/blubird/postgres"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("No .env file AHHHH!!!")
    }
}

func main() {
    // Connect to database
    db, err := postgres.Open()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    rs := postgres.ResortService{DB: db}
    r, err := rs.Resort("sierra")
    if err != nil {
        panic(err)
    }

    fmt.Println(r.Name)
}
