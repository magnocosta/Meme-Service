package types

import (
	"net/http"
	"time"
)

type Input interface {
  GetCustomerID() string    
  GetTokens()     int       
}

type Output interface {
  Input
  GetUpdatedAt()  time.Time 
}

type DB interface {
  Create(input Input) (Output, error)
}

type UseCase interface {
  Execute(input Input) (Output, error)
}

type HTTP interface {
  Handle(w http.ResponseWriter, r *http.Request)
}

