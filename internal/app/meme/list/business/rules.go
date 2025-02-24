package business

import (
	"fmt"
	"meme_service/internal/app/meme/list/types"
)

var urls []string = []string{"a", "b", "c", "d"}

type business struct {
  db types.DB
  service types.Service
}

func (b business) Execute(input types.Input) (types.Outputs, error) {
  if valid, err := validate(input); !valid {
    return nil, err
  }

  token, err := b.db.ConsumerToken(input)
  if err != nil {
    return nil, err
  }

  if token <= 0 {
    return nil, fmt.Errorf("there is no token available")
  }

  return b.service.List(input)
}

func New(db types.DB, service types.Service) types.UseCase {
  return business{
    db: db,
    service: service,
  }
}

