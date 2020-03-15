package postgres

import (
    "database/sql"
    _ "github.com/lib/pq"

    "github.com/d1rtyvans/blubird/app"
)

type ResortService struct {
    DB *sql.DB
}

func (s *ResortService) Resort(uid string) (*app.Resort, error) {
    var r app.Resort

    sqlQuery := `SELECT id, uid, name, lat, lon FROM resorts WHERE uid = $1`
    row := s.DB.QueryRow(sqlQuery, uid)

    if err := row.Scan(&r.Id, &r.Uid, &r.Name, &r.Lat, &r.Lon); err != nil {
        return nil, err
    }

    return &r, nil
}

func (s *ResortService) Resorts() ([]*app.Resort, error) {
    var resorts []*app.Resort

    sqlQuery := `SELECT id, uid, name, lat, lon FROM resorts`
    rows, err := s.DB.Query(sqlQuery)
    defer rows.Close()

    if err != nil {
        return nil, err
    }

    for rows.Next() {
        var r app.Resort
        err := rows.Scan(&r.Id, &r.Uid, &r.Name, &r.Lat, &r.Lon)
        if err != nil {
            return nil, err
        }
        resorts = append(resorts, &r)
    }

    return resorts, nil
}
