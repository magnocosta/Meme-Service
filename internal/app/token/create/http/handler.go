package http

import (
	"encoding/json"
	"meme_service/internal/app/token/create/types"
	"meme_service/internal/shared/errors"
	sharedtypes "meme_service/internal/shared/types"
	"net/http"
)

type request = sharedtypes.Token

type controller struct {
  useCase types.UseCase
}

func (c controller) Handle(w http.ResponseWriter, r *http.Request) {
  var req request
  json.NewDecoder(r.Body).Decode(&req)

  result, err := c.useCase.Execute(req)
  if err != nil {
    errors.BuildError(w, err)
    return
  }

  json.NewEncoder(w).Encode(result)
}

func New(useCase types.UseCase) types.HTTP {
  return controller{
    useCase: useCase,
  }
}
