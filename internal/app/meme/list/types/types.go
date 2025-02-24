package types

import (
	"net/http"
  sharedtypes "meme_service/internal/shared/types"
)

type Input interface {
  GetCoordinate() sharedtypes.Coordinate
  GetQuery() string
  GetCustomerID() string
}

type Output interface {
  GetURL() string
}

type Outputs []Output

type DB interface {
  ConsumerToken(input Input) (int, error)
}

type UseCase interface {
  Execute(input Input) (Outputs, error)
}

type HTTP interface {
  Handle(w http.ResponseWriter, r *http.Request)
}

type Service interface {
  List(input Input) (Outputs, error)
}

type Request struct {
  Coordinate sharedtypes.Coordinate
  CustomerID string
  Query string
}

func (r Request) GetCoordinate() sharedtypes.Coordinate { return r.Coordinate  }
func (r Request) GetQuery()      string                 { return r.Query       }
func (r Request) GetCustomerID() string                 { return r.CustomerID  }
