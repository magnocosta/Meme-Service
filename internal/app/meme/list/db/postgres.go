package db

import (
  _ "embed"
	"database/sql"
	"fmt"

  "meme_service/internal/app/meme/list/types"
)

//go:embed query.sql
var query string

type postgres struct {
  conn *sql.DB
}

func (c postgres) ConsumerToken(input types.Input) (int, error) {
  var result int

  fmt.Println(query)

  err := c.conn.QueryRow(query, input.GetCustomerID()).Scan(
    &result,
  )

  if err != nil {
    return result, fmt.Errorf("error on get token: %v", err)
  }

  return result, nil
}

func New(conn *sql.DB) types.DB {
  return postgres{
    conn: conn,
  }
}
