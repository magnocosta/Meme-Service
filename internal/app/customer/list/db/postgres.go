package db

import (
  _ "embed"
	"database/sql"
	"fmt"

  "meme_service/internal/app/customer/list/types"
	sharedtypes "meme_service/internal/shared/types"
)

//go:embed query.sql
var query string

type postgres struct {
  conn *sql.DB
}

func (c postgres) List() (types.Outputs, error) {
  var result types.Outputs

  rows, err := c.conn.Query(query)
  if err != nil {
    return nil, fmt.Errorf("error on list customers: %v", err)
  }
  defer rows.Close()

  for rows.Next() {
		var c sharedtypes.Customer
		err := rows.Scan(&c.ID, &c.ExternalID, &c.Name, &c.Email, &c.CreatedAt)
		if err != nil {
      return nil, fmt.Errorf("error on scanning customers: %v", err)
		}
		result = append(result, c)
	}

  return result, nil
}

func New(conn *sql.DB) types.DB {
  return postgres{
    conn: conn,
  }
}

