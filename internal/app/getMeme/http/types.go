package http

import "meme_service/internal/shared/types"

type request struct {
  CustomerID string
  Coordinate types.Coordinate
  Query string
}

func (d request) GetCoordinate() types.Coordinate { return d.Coordinate   }
func (d request) GetQuery() string { return d.Query }
func (d request) GetCustomerID() string { return d.CustomerID }

type item struct {
  URL string `json:"url"`
}

type response struct {
  Items []item `json:"items"`
}
