package business

import (
  "meme_service/internal/db"
)

func Execute(input Input) error {
  return db.AddCustomerToken(input)
}
