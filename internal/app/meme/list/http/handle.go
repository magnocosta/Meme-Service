package http

import (
	"encoding/json"
	"meme_service/internal/app/meme/list/types"
	"meme_service/internal/shared/errors"
	sharedtypes "meme_service/internal/shared/types"
	"net/http"
)

type request = sharedtypes.Meme

type controller struct {
  useCase types.UseCase
}

func (c controller) Handle(w http.ResponseWriter, r *http.Request) {
  req := types.Request{
    Coordinate: sharedtypes.Coordinate{
      Lat: sharedtypes.Latitude{ Value: r.URL.Query().Get("lat")},
      Lon: sharedtypes.Longitude{ Value: r.URL.Query().Get("lon")},
    },
    CustomerID: r.Header.Get("X-Meme-Access-Token"),
    Query: r.URL.Query().Get("query"),
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
