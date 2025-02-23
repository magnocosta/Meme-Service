package business

import (
  "meme_service/internal/app/customer/list/types"
)

type business struct {
  db types.DB
}

func (d business) Execute() (types.Outputs, error) {
  outputs, err := d.db.List()
  if err != nil {
    return nil, err
  }
  return outputs, nil
}

func New(db types.DB) types.UseCase {
  return business{
    db: db,
  }
}

