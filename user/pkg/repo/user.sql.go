// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package repo

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :one
insert into users (username, hashed_password, first_name, last_name, birth_day, gender, email, phone_number)
values ($1, $2, $3, $4, $5, $6, $7, $8)
returning id, username, hashed_password, first_name, last_name, birth_day, gender, email, phone_number, updated_at, created_at
`

type CreateUserParams struct {
	Username       sql.NullString `json:"username"`
	HashedPassword string         `json:"hashed_password"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	BirthDay       sql.NullTime   `json:"birth_day"`
	Gender         int16          `json:"gender"`
	Email          string         `json:"email"`
	PhoneNumber    sql.NullString `json:"phone_number"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FirstName,
		arg.LastName,
		arg.BirthDay,
		arg.Gender,
		arg.Email,
		arg.PhoneNumber,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.BirthDay,
		&i.Gender,
		&i.Email,
		&i.PhoneNumber,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getHashedPassword = `-- name: GetHashedPassword :one
select hashed_password
from users
where id = $1
limit 1
`

func (q *Queries) GetHashedPassword(ctx context.Context, id int64) (string, error) {
	row := q.db.QueryRowContext(ctx, getHashedPassword, id)
	var hashed_password string
	err := row.Scan(&hashed_password)
	return hashed_password, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
select id,
       username,
       first_name,
       last_name,
       birth_day,
       gender,
       email,
       phone_number,
       updated_at,
       created_at
from users
where email = $1
limit 1
`

type GetUserByEmailRow struct {
	ID          int64          `json:"id"`
	Username    sql.NullString `json:"username"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	BirthDay    sql.NullTime   `json:"birth_day"`
	Gender      int16          `json:"gender"`
	Email       string         `json:"email"`
	PhoneNumber sql.NullString `json:"phone_number"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	CreatedAt   time.Time      `json:"created_at"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.BirthDay,
		&i.Gender,
		&i.Email,
		&i.PhoneNumber,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
select id,
       username,
       first_name,
       last_name,
       birth_day,
       gender,
       email,
       phone_number,
       updated_at,
       created_at
from users
where id = $1
limit 1
`

type GetUserByIDRow struct {
	ID          int64          `json:"id"`
	Username    sql.NullString `json:"username"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	BirthDay    sql.NullTime   `json:"birth_day"`
	Gender      int16          `json:"gender"`
	Email       string         `json:"email"`
	PhoneNumber sql.NullString `json:"phone_number"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	CreatedAt   time.Time      `json:"created_at"`
}

func (q *Queries) GetUserByID(ctx context.Context, id int64) (GetUserByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.BirthDay,
		&i.Gender,
		&i.Email,
		&i.PhoneNumber,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
select id,
       username,
       first_name,
       last_name,
       birth_day,
       gender,
       email,
       phone_number,
       updated_at,
       created_at
from users
where username = $1
limit 1
`

type GetUserByUsernameRow struct {
	ID          int64          `json:"id"`
	Username    sql.NullString `json:"username"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	BirthDay    sql.NullTime   `json:"birth_day"`
	Gender      int16          `json:"gender"`
	Email       string         `json:"email"`
	PhoneNumber sql.NullString `json:"phone_number"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	CreatedAt   time.Time      `json:"created_at"`
}

func (q *Queries) GetUserByUsername(ctx context.Context, username sql.NullString) (GetUserByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i GetUserByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.BirthDay,
		&i.Gender,
		&i.Email,
		&i.PhoneNumber,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserEmail = `-- name: UpdateUserEmail :one
update users
set email = $2
where id = $1
returning id, username, hashed_password, first_name, last_name, birth_day, gender, email, phone_number, updated_at, created_at
`

type UpdateUserEmailParams struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func (q *Queries) UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserEmail, arg.ID, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.BirthDay,
		&i.Gender,
		&i.Email,
		&i.PhoneNumber,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserInfo = `-- name: UpdateUserInfo :one
update users
set first_name = $2 and last_name = $3 and birth_day = $4 and gender = $5
where id = $1
returning id, username, hashed_password, first_name, last_name, birth_day, gender, email, phone_number, updated_at, created_at
`

type UpdateUserInfoParams struct {
	ID        int64        `json:"id"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	BirthDay  sql.NullTime `json:"birth_day"`
	Gender    int16        `json:"gender"`
}

func (q *Queries) UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserInfo,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.BirthDay,
		arg.Gender,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.BirthDay,
		&i.Gender,
		&i.Email,
		&i.PhoneNumber,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserPassword = `-- name: UpdateUserPassword :one
update users
set hashed_password = $2
where id = $1
returning id, username, hashed_password, first_name, last_name, birth_day, gender, email, phone_number, updated_at, created_at
`

type UpdateUserPasswordParams struct {
	ID             int64  `json:"id"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserPassword, arg.ID, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.BirthDay,
		&i.Gender,
		&i.Email,
		&i.PhoneNumber,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserPhoneNumber = `-- name: UpdateUserPhoneNumber :one
update users
set phone_number = $2
where id = $1
returning id, username, hashed_password, first_name, last_name, birth_day, gender, email, phone_number, updated_at, created_at
`

type UpdateUserPhoneNumberParams struct {
	ID          int64          `json:"id"`
	PhoneNumber sql.NullString `json:"phone_number"`
}

func (q *Queries) UpdateUserPhoneNumber(ctx context.Context, arg UpdateUserPhoneNumberParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserPhoneNumber, arg.ID, arg.PhoneNumber)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.BirthDay,
		&i.Gender,
		&i.Email,
		&i.PhoneNumber,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserUsername = `-- name: UpdateUserUsername :one
update users
set username = $2
where id = $1
returning id, username, hashed_password, first_name, last_name, birth_day, gender, email, phone_number, updated_at, created_at
`

type UpdateUserUsernameParams struct {
	ID       int64          `json:"id"`
	Username sql.NullString `json:"username"`
}

func (q *Queries) UpdateUserUsername(ctx context.Context, arg UpdateUserUsernameParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserUsername, arg.ID, arg.Username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.BirthDay,
		&i.Gender,
		&i.Email,
		&i.PhoneNumber,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}