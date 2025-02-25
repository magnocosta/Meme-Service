package db

import (
  _ "embed"
	"database/sql"
	"fmt"

  "meme_service/internal/app/token/create/types"
	sharedtypes "meme_service/internal/shared/types"
)

//go:embed query.sql
var query string

type postgres struct {
  conn *sql.DB
}

func (c postgres) Create(input types.Input) (types.Output, error) {
  var result sharedtypes.Token

  err := c.conn.QueryRow(query, input.GetCustomerID(), input.GetTokens()).Scan(
    &result.Tokens,
    &result.UpdatedAt,
  )

  if err != nil {
    return nil, fmt.Errorf("error on add token: %v", err)
  }

  return result, nil
}

func NewPostgres(conn *sql.DB) types.DB {
  return postgres{
    conn: conn,
  }
}

