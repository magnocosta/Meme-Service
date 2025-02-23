package addcustomer

import (
	"meme_service/internal/app/customer/create/business"
	"meme_service/internal/app/customer/create/db"
	"meme_service/internal/app/customer/create/types"
	"meme_service/internal/app/customer/create/http"
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
