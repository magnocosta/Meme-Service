package http

import (
	"encoding/json"
	"meme_service/internal/app/addCustomer/types"
	"meme_service/internal/shared/errors"
	sharedtypes "meme_service/internal/shared/types"
	"net/http"
)

type request  = sharedtypes.Customer
type response = sharedtypes.Customer

type CreateCustomerHTTP struct {
  useCase types.CreateCustomerUseCase
}

func (c CreateCustomerHTTP) Handle(w http.ResponseWriter, r *http.Request) {
  var req request
  json.NewDecoder(r.Body).Decode(&req)

  result, err := c.useCase.Execute(req)
  if err != nil {
    errors.BuildError(w, err)
    return
  }

  json.NewEncoder(w).Encode(result)
}

func NewCreateCustomerHTTP(useCase types.CreateCustomerUseCase) types.CreateCustomerHTTP {
  return &CreateCustomerHTTP{
    useCase: useCase,
  }
}
