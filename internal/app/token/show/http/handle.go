package http

import (
	"encoding/json"
	"meme_service/internal/app/token/show/types"
	"meme_service/internal/shared/errors"
	sharedtypes "meme_service/internal/shared/types"
	"meme_service/internal/shared/utils"
	"net/http"
)

type request = sharedtypes.Token

type controller struct {
  useCase types.UseCase
}

func (c controller) Handle(w http.ResponseWriter, r *http.Request) {
  vars := utils.GetVars(r)
  req := request{
    CustomerID: vars["customer_id"],
  }

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
