package db

import (
  _ "embed"
	"database/sql"
	"fmt"

  "meme_service/internal/app/meme/list/types"
  sharedtypes "meme_service/internal/shared/types"
)

//go:embed query.sql
var query string

type postgres struct {
  conn *sql.DB
}

func (c postgres) ConsumerToken(input types.Input) (types.ConsumerTokenOuput, error) {
  var result sharedtypes.Token
  result.CustomerID = input.GetCustomerID()

  err := c.conn.QueryRow(query, input.GetCustomerID()).Scan(
    &result.Tokens,
  )

  if err != nil {
    return result, fmt.Errorf("error on get token: %v", err)
  }

  if result.GetTokens() <= 0 {
    return nil, fmt.Errorf("there is no token available")
  }

  return result, nil
}

func NewPostgres(conn *sql.DB) types.DB {
  return postgres{
    conn: conn,
  }
}
