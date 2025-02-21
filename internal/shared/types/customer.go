package types

type Customer struct {
  ID    string `json:"id"`
  Name  string `json:"name"`
  Email string `json:"email"`
}

func (c Customer) GetName()  string { return c.Name  }
func (c Customer) GetEmail() string { return c.Email }
func (c Customer) GetID()    string { return c.ID    }
