// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package postgres

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type BaseProduct struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Customer struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Address   string
	Phone     string
	Email     string
	Latitude  float32
	Longitude float32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Inventory struct {
	ID              uuid.UUID
	UserID          uuid.UUID
	ProductID       uuid.UUID
	Quantity        int32
	UnitOfMeasureID uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Product struct {
	ID            uuid.UUID
	Name          string
	BaseProductID uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Role struct {
	ID        uuid.UUID
	Name      sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UnitsOfMeasure struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Address   string
	Phone     string
	Email     string
	Latitude  float32
	Longitude float32
	RoleID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
