package db

import (
  _ "embed"
	"database/sql"
	"fmt"
	"meme_service/internal/app/customer/create/types"
	sharedtypes "meme_service/internal/shared/types"
)

//go:embed query.sql
var query string

type postgresDB struct {
  conn *sql.DB
}

func (c postgresDB) Create(customer types.Input) (types.Output, error) {
  var result sharedtypes.Customer

  err := c.conn.QueryRow(query, customer.GetName(), customer.GetEmail()).Scan(
    &result.ID,
    &result.ExternalID,
		&result.Name,
		&result.Email,
    &result.CreatedAt,
  )

  if err != nil {
    return nil, fmt.Errorf("error on create customer: %v", err)
  }

  return result, nil
}

func NewCreateCustomerDB(conn *sql.DB) types.CreateCustomerDB {
  return postgresDB{
    conn: conn,
  }
}
