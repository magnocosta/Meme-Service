package http

import (
	"encoding/json"
	"meme_service/internal/app/getMeme/business"
	"meme_service/internal/shared/types"
	"net/http"
)

func buildRequest(r *http.Request) request {
  return request{
    Coordinate: types.Coordinate{
      Lat: types.Latitude{ Value: r.URL.Query().Get("lat")},
      Lon: types.Longitude{ Value: r.URL.Query().Get("lon")},
    },
    Query: r.URL.Query().Get("query"),
  }
}

func buildResponse(w http.ResponseWriter, result business.Output) {
  response := response{}
  for _, v := range result.GetItems() {
    response.Items = append(response.Items, item{ URL: v.GetURL()})
  }

  json.NewEncoder(w).Encode(response)
}
