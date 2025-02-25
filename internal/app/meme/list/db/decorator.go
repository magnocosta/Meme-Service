package db

import (
	"database/sql"
	"meme_service/internal/app/meme/list/types"
	shareddb "meme_service/internal/db"
)

type decorator struct {
  influxdb influxdb
  postgres postgres
}

func (i decorator) ConsumerToken(input types.Input) (int, error) {
  result, err := i.postgres.ConsumerToken(input)
  if err != nil {
    return 0, err
  }

  if result <= 0 {
    _, err = i.influxdb.ConsumerToken(input)
  }

  return result, err
}

func New(postgresConn *sql.DB, influxDbConn shareddb.InfluxDBConnection) types.DB {
  return decorator{
    postgres: postgres{
      conn: postgresConn,
    },
    influxdb: influxdb{
      conn: influxDbConn,
    },
  } 
}
