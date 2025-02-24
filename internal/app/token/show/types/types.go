package types

import (
	"net/http"
	"time"
)

type Input interface {
  GetCustomerID() string    
}

type Output interface {
  Input
  GetTokens()     int       
  GetUpdatedAt()  time.Time 
}

type DB interface {
  Get(input Input) (Output, error)
}

type UseCase interface {
  Execute(input Input) (Output, error)
}

type HTTP interface {
  Handle(w http.ResponseWriter, r *http.Request)
}

