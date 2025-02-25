package create

import (
	"meme_service/internal/app/token/create/business"
	"meme_service/internal/app/token/create/db"
	"meme_service/internal/app/token/create/types"
	"meme_service/internal/app/token/create/http"
	shareddb "meme_service/internal/db"
	"database/sql"
)

func NewUseCase(postgresConn *sql.DB, influxDbConn shareddb.InfluxDBConnection) types.UseCase {
  db := db.New(postgresConn, influxDbConn)
  useCase := business.New(db)
  return useCase
}

func NewHTTPHandler(postgresConn *sql.DB, influxDbConn shareddb.InfluxDBConnection) types.HTTP {
  useCase := NewUseCase(postgresConn, influxDbConn)
  http := http.New(useCase)
  return http
}
