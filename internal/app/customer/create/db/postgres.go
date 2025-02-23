package db

import (
  _ "embed"
	"database/sql"
	"fmt"
	"meme_service/internal/app/customer/create/types"
	sharedtypes "meme_service/internal/shared/types"
)

//go:embed customer.sql
var insertCustomerQuery string

//go:embed customer_tokens.sql
var insertCustomerTokenQuery string


type postgresDB struct {
  conn *sql.DB
}

func (c postgresDB) Create(customer types.Input) (types.Output, error) {
  var result sharedtypes.Customer

  tx, err := c.conn.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}

  err = tx.QueryRow(insertCustomerQuery, customer.GetName(), customer.GetEmail()).Scan(
    &result.ID,
    &result.ExternalID,
		&result.Name,
		&result.Email,
    &result.CreatedAt,
  )

  if err != nil {
    tx.Rollback()
    return nil, fmt.Errorf("error on create customer: %v", err)
  }

  _, err = tx.Exec(insertCustomerTokenQuery, result.ID)
  if err != nil {
    tx.Rollback()
    return nil, fmt.Errorf("error on create customer_tokens: %v", err)
  }

  err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

  return result, nil
}

func NewCreateCustomerDB(conn *sql.DB) types.CreateCustomerDB {
  return postgresDB{
    conn: conn,
  }
}
