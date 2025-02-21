package http

type request struct {
  CustomerID string `json:"customer_id"`
  Tokens int `json:"tokens"`
}

func (r request) GetCustomerID() string { return r.CustomerID }
func (r request) GetTokens() int { return r.Tokens }
