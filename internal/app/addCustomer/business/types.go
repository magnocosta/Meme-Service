package business

type Input interface {
  GetName() string
  GetEmail() string
}

type Output interface {
  Input
  GetID() string
}
