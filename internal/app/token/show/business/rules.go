package business

import (
  "meme_service/internal/app/token/show/types"
)

type business struct {
  db types.DB
}

func (b business) Execute(input types.Input) (types.Output, error) {
  return b.db.Get(input)
}

func New(db types.DB) types.UseCase {
  return business{
    db: db,
  }
}
