package db

type CustomerInput interface {
  GetName()  string
  GetEmail() string
}

type CustomerOuput interface {
  CustomerInput
  GetID() string
}

type CustomerTokenInput interface {
  GetCustomerID() string
  GetTokens() int 
}

type ConsumerTokenInput interface {
  GetCustomerID() string
}
