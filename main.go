package main

import (
    "log"
    "fmt"

    "github.com/joho/godotenv"

    // _ "github.com/d1rtyvans/blubird/app"    // TODO: May not need
    "github.com/d1rtyvans/blubird/config" // TODO: Set this up with env variables
    "github.com/d1rtyvans/blubird/postgres"
    "github.com/d1rtyvans/blubird/http"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("No .env file AHHHH!!!")
    }
}

func main() {
    conf := config.New()
    // Connect to database
    db, err := postgres.Open()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    rs := postgres.ResortService{DB: db}
    r, err := rs.Resort("sierra")
    if err != nil {
        log.Fatal(err)
    }

    wc := http.WeatherClient{ApiKey: conf.DarkSky.ApiKey}
    resp, err := wc.DayForecasts(r.Lat, r.Lon)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(resp))
    // fs := postgres.ForecastService{DB: db}
}
