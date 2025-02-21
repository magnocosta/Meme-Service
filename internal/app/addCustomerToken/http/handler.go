package http

import (
	"encoding/json"
	"meme_service/internal/app/addCustomerToken/business"
	"meme_service/internal/shared/errors"
	"net/http"
)

func AddCustomerToken(w http.ResponseWriter, r *http.Request) {
  var req request
  json.NewDecoder(r.Body).Decode(&req)

  err := business.Execute(req)
  if err != nil {
    errors.BuildError(w, err)
    return
  }
}
