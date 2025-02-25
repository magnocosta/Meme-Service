package db

import (
	"context"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxDBConnection interface {
  GetOrd() string
  GetBucket() string
  GetClient() influxdb2.Client
}

type influxDBConfig struct {
  Org    string
  Bucket string
  client influxdb2.Client
}

func (i influxDBConfig) GetOrd() string { return i.Org }
func (i influxDBConfig) GetBucket() string { return i.Bucket }
func (i influxDBConfig) GetClient() influxdb2.Client { return i.client }

func NewInfluxDB(endpoint string, secret string, org string, bucket string) (InfluxDBConnection, error) {
  client := influxdb2.NewClient(endpoint, secret)
	defer client.Close()

  _, err := client.Ping(context.Background())
  if err != nil {
    return nil, err
  }

  config := influxDBConfig{
    Org: org,
    Bucket: bucket,
    client: client,
  }

  return config, nil
}
