package list

import (
	"meme_service/internal/app/customer/list/business"
	"meme_service/internal/app/customer/list/db"
	"meme_service/internal/app/customer/list/types"
	"meme_service/internal/app/customer/list/http"
	"database/sql"
)

func NewUseCase(conn *sql.DB) types.UseCase {
  db := db.New(conn)
  useCase := business.New(db)
  return useCase
}

func NewHTTPHandler(conn *sql.DB) types.HTTP {
  useCase := NewUseCase(conn)
  http := http.New(useCase)
  return http
}
