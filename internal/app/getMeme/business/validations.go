package business

import "fmt"

func validate(input Input) (bool, error) {
  var valid bool
  var err error

  if input.GetCoordinate().IsEmpty() {
    if input.GetQuery() == "" {
      return false, fmt.Errorf("Invalid input: query or coordinate must be filled")
    }
    return true, nil
  } else {
    valid, err = input.GetCoordinate().IsValid()
  }

  return valid, err
}
