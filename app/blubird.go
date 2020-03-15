package app

type Resort struct {
    Id   int
    Uid  string
    Name string
    Lat  float64
    Lon  float64
}

type ResortService interface {
    Resort(uid string) (*Resort, error)
    Resorts() ([]*Resort, error)
}
