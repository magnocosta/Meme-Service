package db

import (
	"database/sql"
	"fmt"
)

type Config struct {
  Host string
  Port int
  User string
  Password string
  Dbname string
  Driver string
}

func Connect(config *Config) (*sql.DB, error) {
  connURL := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
  dbInfo := fmt.Sprintf(connURL, config.Host, config.Port, config.User, config.Password, config.Dbname)

  db, err := sql.Open(config.Driver, dbInfo)
  if err != nil {
    return nil, err
  }

  err = db.Ping()
  if err != nil {
    return nil, err
  }

  return db, nil
}


