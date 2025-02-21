package db

type CustomerInput interface {
  GetName()  string
  GetEmail() string
}

type CustomerOuput interface {
  CustomerInput
  GetID() string
}
