package http

import (
	getMemes "meme_service/internal/app/getMeme/http"
	addCustomer "meme_service/internal/app/addCustomer/http"
	listCustomers "meme_service/internal/app/listCustomers/http"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
  router := mux.NewRouter()

  router.HandleFunc("/v1/memes", getMemes.ListMemesHandle).Methods("Get")
  router.HandleFunc("/v1/customers", addCustomer.CreateCustomerHandle).Methods("Post")
  router.HandleFunc("/v1/customers", listCustomers.ListCustomersHandle).Methods("Get")

  http.ListenAndServe(":8080", router)
}
