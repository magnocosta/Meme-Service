package types

import (
	"fmt"
	"strconv"
)

type Longitude struct {
  Value string
}

func (l Longitude) IsValid() (bool, error) {
  value, err := strconv.ParseFloat(l.Value, 64)
  if err != nil {
    return false, fmt.Errorf("Longitude is not a number")
  }

  if (value < -180 || value > 180) {
    return false, fmt.Errorf("Longitude invalid: it must be between -180 and 180")
  }

  return true, nil
}

func (l Longitude) IsEmpty() bool {
  return l.Value == ""
}
