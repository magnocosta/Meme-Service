package business

import (
	"meme_service/internal/app/meme/list/types"
)

var urls []string = []string{"a", "b", "c", "d"}

type business struct {
  db types.DB
  service types.Service
  publisher types.Publisher
}

func (b business) Execute(input types.Input) (types.Outputs, error) {
  if valid, err := validate(input); !valid {
    return nil, err
  }

  result, err := b.db.ConsumerToken(input)
  if err != nil {
    return nil, err
  }

  // TODO - Magno Costa: This should be optional per customer. 
  // When the customer opens the real-time app, we should enable this behavior.
  err = b.publisher.Publish(result)
  if err != nil {
    return nil, err
  }

  return b.service.List(input)
}

func New(db types.DB, service types.Service, publisher types.Publisher) types.UseCase {
  return business{
    db: db,
    service: service,
    publisher: publisher,
  }
}

