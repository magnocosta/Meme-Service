package db

import (
	"context"
	"fmt"
  "meme_service/internal/app/token/create/types"
	"meme_service/internal/db"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type influxdb struct {
  conn db.InfluxDBConnection
}

func (i influxdb) Create(input types.Input) (types.Output, error) {
  writeAPI := i.conn.GetClient().WriteAPIBlocking(i.conn.GetOrd(), i.conn.GetBucket())

	tags := map[string]string{
		"customer_id": fmt.Sprintf("%v", input.GetCustomerID()),
    "transaction_type": "replenishment",
	}
	fields := map[string]interface{}{
		"value": input.GetTokens(),
	}

	point := influxdb2.NewPoint("customer_token_balance", tags, fields, time.Now())
	err := writeAPI.WritePoint(context.Background(), point)
	if err != nil {
		return nil, fmt.Errorf("error writing to InfluxDB: %v", err)
	}

  return nil, nil
}

func NewInfluxDB(conn db.InfluxDBConnection) types.DB {
  return influxdb{
    conn: conn,
  }
}
