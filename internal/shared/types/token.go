package types

import "time"

type Token struct {
  CustomerID string    `json:"customer_id,omitempty"` 
  Tokens     int       `json:"tokens"`
  UpdatedAt  time.Time `json:"updated_at"`
}

func (c Token) GetCustomerID() string    { return c.CustomerID }
func (c Token) GetTokens()     int       { return c.Tokens     }
func (c Token) GetUpdatedAt()  time.Time { return c.UpdatedAt  }
