package types

import "time"

type Customer struct {
  ID         string    `json:"-"` 
  ExternalID string    `json:"id"`
  Name       string    `json:"name"`
  Email      string    `json:"email"`
  CreatedAt  time.Time `json:"created_at"`
  Tokens     int       `json:"tokens"`
}

func (c Customer) GetName()  string { return c.Name  }
func (c Customer) GetEmail() string { return c.Email }
func (c Customer) GetID()    string { return c.ID    }
