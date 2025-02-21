package http

import (
	"meme_service/internal/app/getMeme/business"
	"meme_service/internal/shared/errors"
	"net/http"
)

func ListMemesHandle(w http.ResponseWriter, r *http.Request) {
  req := buildRequest(r)

  result, err := business.Execute(req)
  if err != nil {
    errors.BuildError(w, err)
    return
  }

  buildResponse(w, result)
}
