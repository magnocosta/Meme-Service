package create

import (
	"meme_service/internal/app/token/create/business"
	"meme_service/internal/app/token/create/db"
	"meme_service/internal/app/token/create/types"
	"meme_service/internal/app/token/create/http"
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
