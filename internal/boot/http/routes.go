package http

import (
	"database/sql"
	createCustomer "meme_service/internal/app/customer/create"
	listCustomer "meme_service/internal/app/customer/list"
	createToken "meme_service/internal/app/token/create"
	showToken "meme_service/internal/app/token/show"
	"meme_service/internal/db"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Init() {
  router := mux.NewRouter()
  conn := createDB()

  createCustomerHandler := createCustomer.NewHTTPHandler(conn)
  router.HandleFunc("/v1/customers", createCustomerHandler.Handle).Methods("Post")

  listCustomerHandler := listCustomer.NewHTTPHandler(conn)
  router.HandleFunc("/v1/customers", listCustomerHandler.Handle).Methods("Get")

  createTokenHandler := createToken.NewHTTPHandler(conn)
  router.HandleFunc("/v1/customers/{customer_id}/tokens", createTokenHandler.Handle).Methods("Post")

  showTokenHandler := showToken.NewHTTPHandler(conn)
  router.HandleFunc("/v1/customers/{customer_id}/tokens", showTokenHandler.Handle).Methods("Get")

  // router.HandleFunc("/v1/memes", getMemes.ListMemesHandle).Methods("Get")
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
