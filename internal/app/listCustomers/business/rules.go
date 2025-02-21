package business

import (
  "meme_service/internal/db"
)

func Execute() (Output, error) {
  items, err := db.ListCustomers()
  if err != nil {
    return nil, err
  }

  output := output{}
  for _, v := range items {
    output.items = append(output.items, v)
  }

  return output, nil
}
