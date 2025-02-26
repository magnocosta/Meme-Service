package types

type Coordinate struct {
  Lon Longitude // TODO - Magno Costa: I would like to add more custom type like this
  Lat Latitude  // Instead of apply this for all data collected I've focused the busines rules.
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
