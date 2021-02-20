// Code generated by sqlc. DO NOT EDIT.

package repo

import (
	"database/sql"
	"time"
)

type Gender struct {
	ID    int16  `json:"id"`
	Title string `json:"title"`
}

type User struct {
	ID             int64          `json:"id"`
	Email          string         `json:"email"`
	HashedPassword string         `json:"hashed_password"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	BirthDay       sql.NullTime   `json:"birth_day"`
	Gender         int16          `json:"gender"`
	PhoneNumber    sql.NullString `json:"phone_number"`
	UpdatedAt      sql.NullTime   `json:"updated_at"`
	DeletedAt      sql.NullTime   `json:"deleted_at"`
	CreatedAt      time.Time      `json:"created_at"`
}
