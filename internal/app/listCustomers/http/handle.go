package http

import (
  "net/http" 
  "encoding/json"
  "meme_service/internal/app/listCustomers/business"
	"meme_service/internal/shared/errors"
	"meme_service/internal/shared/types"
)

type response = []types.Customer

func ListCustomersHandle(w http.ResponseWriter, r *http.Request) {
  result, err := business.Execute()
  if err != nil {
    errors.BuildError(w, err)
    return
  }

  resp := response{}
  for _, v := range result.GetItems() {
    customer := types.Customer{
      ID: v.GetID(),
      Name: v.GetName(),
      Email: v.GetEmail(),
    }

    resp = append(resp, customer)
  }

  json.NewEncoder(w).Encode(resp)
}
