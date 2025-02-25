package list

import (
	"database/sql"
	"meme_service/internal/app/meme/list/business"
	"meme_service/internal/app/meme/list/db"
	"meme_service/internal/app/meme/list/http"
	"meme_service/internal/app/meme/list/service"
	"meme_service/internal/app/meme/list/types"
	shareddb "meme_service/internal/db"
)

func NewUseCase(postgresConn *sql.DB, influxDbConn shareddb.InfluxDBConnection) types.UseCase {
  service := service.New()
  db := db.New(postgresConn, influxDbConn)
  useCase := business.New(db, service)
  return useCase
}

func NewHTTPHandler(postgresConn *sql.DB, influxDbConn shareddb.InfluxDBConnection) types.HTTP {
  useCase := NewUseCase(postgresConn, influxDbConn)
  http := http.New(useCase)
  return http
}
