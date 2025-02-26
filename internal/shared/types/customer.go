package types

import "time"

type Customer struct {
  ID         string    `json:"-"` 
  ExternalID string    `json:"id"` // TODO - Magno Costa: I would like to only expose this field externaly
  Name       string    `json:"name"`
  Email      string    `json:"email"`
  CreatedAt  time.Time `json:"created_at"`
  // TODO - Magno Costa: I would like to add an access_token with JWT.
}

func (c Customer) GetName()  string { return c.Name  }
func (c Customer) GetEmail() string { return c.Email }
func (c Customer) GetID()    string { return c.ID    }
