package business

import (
  "meme_service/internal/app/token/create/types"
)

type business struct {
  db types.DB
}

func (b business) Execute(input types.Input) (types.Output, error) {
  return b.db.Create(input)
}

func New(db types.DB) types.UseCase {
  return business{
    db: db,
  }
}
