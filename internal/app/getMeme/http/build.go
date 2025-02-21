package http

import (
	"encoding/json"
	"meme_service/internal/app/getMeme/business"
	"net/http"
)

func buildRequest(r *http.Request) request {
  return request{
    Lat: r.URL.Query().Get("lat"),
    Lon: r.URL.Query().Get("lon"),
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

func buildError(w http.ResponseWriter, err error) {
  responseError := responseError{
    message: err.Error(),
  }
  json.NewEncoder(w).Encode(responseError)
}
