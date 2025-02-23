package types

import "net/http"

type Output interface {
  GetID() string
  GetName() string
  GetEmail() string
}

type Outputs []Output

type DB interface {
  List() (Outputs, error)
}

type UseCase interface {
  Execute() (Outputs, error)
}

type HTTP interface {
  Handle(w http.ResponseWriter, r *http.Request)
}

