package http

import (
	getMemes "meme_service/internal/app/getMeme/http"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
  router := mux.NewRouter()
  router.HandleFunc("/memes", getMemes.ListMemesHandle).Methods("Get")
  http.ListenAndServe(":8080", router)
}
