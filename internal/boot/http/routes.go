package http

import (
	addCustomer "meme_service/internal/app/addCustomer"
	"meme_service/internal/db"
	"net/http"
	"database/sql"

	"github.com/gorilla/mux"
  _ "github.com/lib/pq"
)

func Init() {
  router := mux.NewRouter()
  conn := createDB()

  createCustomerHandler := addCustomer.NewHTTPHandler(conn)
  router.HandleFunc("/v1/customers", createCustomerHandler.Handle).Methods("Post")

  // router.HandleFunc("/v1/memes", getMemes.ListMemesHandle).Methods("Get")
  // router.HandleFunc("/v1/customers", listCustomers.ListCustomersHandle).Methods("Get")
  // router.HandleFunc("/v1/tokens", addCustomerToken.AddCustomerToken).Methods("Post")

  http.ListenAndServe(":8080", router)
}

func createDB() *sql.DB {
  config := &db.Config{
    Host: "localhost",
    Port: 5432,
    User: "meme",
    Password: "meme12",
    Dbname: "meme_db",
    Driver: "postgres",
  }
  conn, err := db.Connect(config)
  if err != nil {
    panic(err)
  }
  return conn
}
