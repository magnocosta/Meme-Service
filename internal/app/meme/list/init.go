package list

import (
	"database/sql"
	"meme_service/internal/app/meme/list/business"
	"meme_service/internal/app/meme/list/db"
	"meme_service/internal/app/meme/list/http"
	"meme_service/internal/app/meme/list/service/meme"
	"meme_service/internal/app/meme/list/service/publisher"
	"meme_service/internal/app/meme/list/types"
  "github.com/aws/aws-sdk-go-v2/service/sns"
	shareddb "meme_service/internal/db"
)

func NewUseCase(postgresConn *sql.DB, influxDbConn shareddb.InfluxDBConnection, snsClient *sns.Client) types.UseCase {
  service := service.New()
  publisher :=  publisher.New(snsClient)
  db := db.New(postgresConn, influxDbConn)
  useCase := business.New(db, service, publisher)
  return useCase
}

func NewHTTPHandler(postgresConn *sql.DB, influxDbConn shareddb.InfluxDBConnection, snsClient *sns.Client) types.HTTP {
  useCase := NewUseCase(postgresConn, influxDbConn, snsClient)
  http := http.New(useCase)
  return http
}
