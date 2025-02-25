package db

import (
	"database/sql"
  "meme_service/internal/app/token/create/types"
	shareddb "meme_service/internal/db"
)

type decorator struct {
  influxdb influxdb
  postgres postgres
}

func (i decorator) Create(input types.Input) (types.Output, error) {
  result, err := i.postgres.Create(input)
  if err != nil {
    return nil, err
  }

  _, err = i.influxdb.Create(input)

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
