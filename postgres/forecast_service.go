package postgres

import (
    "database/sql"
    _ "github.com/lib/pq"

    "github.com/d1rtyvans/blubird/app"
)

type ForecastService struct {
    DB *sql.DB
}

func (s *ForecastService) UpsertAll(forecasts []*app.Forecast) (error) {
    // TODO:
    return nil
}

func (s *ForecastService) CurrentForResort(resortId int) ([]*app.Forecast, error) {
    var forecasts []*app.Forecast

    sqlQuery := `
    SELECT id, date, weather_data
    FROM forecasts
    WHERE resort_id = $1 AND date >= NOW()::date`
    rows, err := s.DB.Query(sqlQuery, resortId)
    defer rows.Close()

    if err != nil {
        return nil, err
    }

    for rows.Next() {
        var f app.Forecast
        err := rows.Scan(&f.Id, &f.Date, &f.WeatherJson)
        if err != nil {
            return nil, err
        }
        forecasts = append(forecasts, &f)
    }

    return forecasts, nil
}
