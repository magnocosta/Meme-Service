package types

type Coordinate struct {
  Lon Longitude
  Lat Latitude
}

func (c Coordinate) IsValid() (bool, error) {
  var valid bool
  var err error

  valid, err = c.Lon.IsValid()
  if !valid {
    return valid, err;
  }

  return c.Lat.IsValid()
}

func (c Coordinate) IsEmpty() bool {
  return c.Lat.IsEmpty() && c.Lon.IsEmpty()
}
