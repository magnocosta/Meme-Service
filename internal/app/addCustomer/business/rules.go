package business

import (
  "meme_service/internal/db"
)

func Execute(input Input) (Output, error) {
  result, err := db.AddCustomer(input)
  return result, err
}
