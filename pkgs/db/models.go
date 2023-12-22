// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"database/sql"

	"idea2/graph/model"
)

type Device struct {
	ID          int64
	Title       string
	Address     string
	ConnectType string
	AuthID      sql.NullInt64
	Note        string
	CreatedAt   model.Timestamp
	UpdatedAt   *model.Timestamp
}

type DeviceAuth struct {
	ID         int64
	Title      string
	Type       string
	Username   *model.NullString
	Password   *model.NullString
	PrivateKey *model.NullString
	CreatedAt  model.Timestamp
	UpdatedAt  *model.Timestamp
}
