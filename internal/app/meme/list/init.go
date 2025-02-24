package list

import (
	"database/sql"
	"meme_service/internal/app/meme/list/business"
	"meme_service/internal/app/meme/list/db"
	"meme_service/internal/app/meme/list/http"
	"meme_service/internal/app/meme/list/service"
	"meme_service/internal/app/meme/list/types"
)

func NewUseCase(conn *sql.DB) types.UseCase {
  service := service.New()
  db := db.New(conn)
  useCase := business.New(db, service)
  return useCase
}

func NewHTTPHandler(conn *sql.DB) types.HTTP {
  useCase := NewUseCase(conn)
  http := http.New(useCase)
  return http
}
