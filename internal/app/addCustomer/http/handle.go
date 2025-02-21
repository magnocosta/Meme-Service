package http

import (
  "net/http" 
  "encoding/json"
  "meme_service/internal/app/addCustomer/business"
	"meme_service/internal/shared/errors"
	"meme_service/internal/shared/types"
)

type request  = types.Customer
type response = types.Customer

func CreateCustomerHandle(w http.ResponseWriter, r *http.Request) {
  var req request
  json.NewDecoder(r.Body).Decode(&req)

  result, err := business.Execute(req)
  if err != nil {
    errors.BuildError(w, err)
    return
  }

  resp := response{
    ID: result.GetID(),
    Name: result.GetName(),
    Email: result.GetEmail(),
  }

  json.NewEncoder(w).Encode(resp)
}
