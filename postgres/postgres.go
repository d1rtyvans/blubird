package postgres

import (
    "fmt"
    "database/sql"

    "github.com/d1rtyvans/blubird/app"
    _ "github.com/lib/pq"
)

const (
    dbHost = "localhost"
    dbPort = 5432
    dbUser = "postgres"
    dbName = "p4p_development"
)

// TODO: Dependency inject environment variables here, similar to 
// https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
func Open() (*sql.DB, error) {
    psqlInfo := fmt.Sprintf(
        "host=%s port=%d user=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbName)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    fmt.Println("Successfully connected to DB!")

    return db, nil
}

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
