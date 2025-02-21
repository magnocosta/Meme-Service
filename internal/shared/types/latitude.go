package types

import (
	"fmt"
	"strconv"
)

type Latitude struct {
  Value string
}

func (l Latitude) IsValid() (bool, error) {
  value, err := strconv.ParseFloat(l.Value, 64)
  if err != nil {
    return false, fmt.Errorf("Latitude is not a number")
  }

  if (value < -90 || value > 90) {
    return false, fmt.Errorf("Latitude invalid: it must be between -90 and 90")
  }

  return true, nil
}

func (l Latitude) IsEmpty() bool {
  return l.Value == ""
}
