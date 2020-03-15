package http

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

const (
    darkSkyUrl = "https://api.darksky.net"
)

// TODO: Dependency inject http?
type WeatherClient struct {
    ApiKey string
}

func (c *WeatherClient) DayForecasts(lat float64, lon float64) ([]byte, error) {
    resp, err := http.Get(c.forecastUrl(lat, lon))
    defer resp.Body.Close()

    if err != nil {
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        return nil, err
    }

    return body, nil
}

func (c *WeatherClient) forecastUrl(lat float64, lon float64) string {
    queryParams := fmt.Sprintf("exclude=currently,minutely,hourly,flags")
    url := fmt.Sprintf("%s/forecast/%s/%g,%g?%s", darkSkyUrl, c.ApiKey, lat, lon, queryParams)

    return url
}
