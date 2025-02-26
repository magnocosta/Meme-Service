package http

import (
	"database/sql"
	createCustomer "meme_service/internal/app/customer/create"
	listCustomer "meme_service/internal/app/customer/list"
	listMeme "meme_service/internal/app/meme/list"
	createToken "meme_service/internal/app/token/create"
	showToken "meme_service/internal/app/token/show"
	internalaws "meme_service/internal/boot/aws"
	"meme_service/internal/db"
	"net/http"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Init() {
  router := mux.NewRouter()
  postgresDB := createDB()
  influxDB := createInfluxDB()
  snsClient := createSNSClient()

  createCustomerHandler := createCustomer.NewHTTPHandler(postgresDB)
  router.HandleFunc("/v1/customers", createCustomerHandler.Handle).Methods("Post")

  listCustomerHandler := listCustomer.NewHTTPHandler(postgresDB)
  router.HandleFunc("/v1/customers", listCustomerHandler.Handle).Methods("Get")

  createTokenHandler := createToken.NewHTTPHandler(postgresDB, influxDB)
  router.HandleFunc("/v1/customers/{customer_id}/tokens", createTokenHandler.Handle).Methods("Post")

  showTokenHandler := showToken.NewHTTPHandler(postgresDB)
  router.HandleFunc("/v1/customers/{customer_id}/tokens", showTokenHandler.Handle).Methods("Get")

  listMemeHandler := listMeme.NewHTTPHandler(postgresDB, influxDB, snsClient)
  router.HandleFunc("/v1/memes", listMemeHandler.Handle).Methods("Get")

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

func createInfluxDB() db.InfluxDBConnection {
  conn, err := db.NewInfluxDB("http://localhost:8086", "secret_token_meme_db", "prix", "meme_db")
  if err != nil {
    panic(err)
  }

  return conn
}

func createSNSClient() *sns.Client {
  config := internalaws.Config{
    URL: "http://localhost:4566",
    Region: "us-east-1",
  }

  client, err := internalaws.New(&config)
  if err != nil {
    panic(err)
  }
  return client
}
