package business

import (
	"meme_service/internal/app/addCustomer/types"
)

type UseCase struct {
  db types.CreateCustomerDB
}

func (c UseCase) Execute(input types.Input) (types.Output, error) {


  output, err := c.db.Create(input)
  return output, err
}

func NewUseCase(db types.CreateCustomerDB) *UseCase {
  return &UseCase{
    db: db,
  }
}

