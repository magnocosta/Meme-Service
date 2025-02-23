package addcustomer

import (
	"meme_service/internal/app/addCustomer/business"
	"meme_service/internal/app/addCustomer/db"
	"meme_service/internal/app/addCustomer/types"
	"meme_service/internal/app/addCustomer/http"
	"database/sql"
)

func NewUseCase(conn *sql.DB) types.CreateCustomerUseCase {
  db := db.NewCreateCustomerDB(conn)
  return business.NewUseCase(db)
}

func NewHTTPHandler(conn *sql.DB) types.CreateCustomerHTTP {
  useCase := NewUseCase(conn)
  return http.NewCreateCustomerHTTP(useCase)
}
