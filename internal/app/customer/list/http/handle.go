package http

import (
  "net/http" 
  "encoding/json"
  "meme_service/internal/app/customer/list/types"
	"meme_service/internal/shared/errors"
)

type controller struct {
  useCase types.UseCase
}

func (c controller) Handle(w http.ResponseWriter, r *http.Request) {
  result, err := c.useCase.Execute()
  if err != nil {
    errors.BuildError(w, err)
    return
  }
  json.NewEncoder(w).Encode(result)
}

func New(useCase types.UseCase) types.HTTP {
  return &controller{
    useCase: useCase,
  }
}
