package show

import (
	"meme_service/internal/app/token/show/business"
	"meme_service/internal/app/token/show/db"
	"meme_service/internal/app/token/show/types"
	"meme_service/internal/app/token/show/http"
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
