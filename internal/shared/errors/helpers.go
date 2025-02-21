package errors

import (
  "net/http"
  "encoding/json"
)

func BuildError(w http.ResponseWriter, err error) {
  responseError := responseError{
    Message: err.Error(),
  }
  json.NewEncoder(w).Encode(responseError)
}
