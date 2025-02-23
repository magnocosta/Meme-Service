package types

import "net/http"

type Input interface {
  GetName() string
  GetEmail() string
}

type Output interface {
  Input
  GetID() string
}

type CreateCustomerDB interface {
  Create(input Input) (Output, error)
}

type CreateCustomerUseCase interface {
  Execute(input Input) (Output, error)
}

type CreateCustomerHTTP interface {
  Handle(w http.ResponseWriter, r *http.Request)
}

