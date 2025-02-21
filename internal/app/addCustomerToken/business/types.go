package business

type Input interface {
  GetCustomerID() string
  GetTokens() int
}
