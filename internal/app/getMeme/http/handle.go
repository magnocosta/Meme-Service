package http

import (
	"meme_service/internal/app/getMeme/business"
	"net/http"
)

func ListMemesHandle(w http.ResponseWriter, r *http.Request) {
  req := buildRequest(r)

  result, err := business.Execute(req)
  if err != nil {
    buildError(w, err)
    return
  }

  buildResponse(w, result)
}
