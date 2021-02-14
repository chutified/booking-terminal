// Code generated by sqlc. DO NOT EDIT.

package repo

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateGender(ctx context.Context, title string) (Gender, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteGender(ctx context.Context, id int16) error
	DeleteUserPermanent(ctx context.Context, id int64) error
	DeleteUserSoft(ctx context.Context, id int64) error
	GetGender(ctx context.Context, arg GetGenderParams) (Gender, error)
	GetHashedPassword(ctx context.Context, id int64) (string, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	GetUserByUsername(ctx context.Context, username sql.NullString) (User, error)
	ListGenders(ctx context.Context) ([]Gender, error)
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (User, error)
	UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) (User, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error)
	UpdateUserPhoneNumber(ctx context.Context, arg UpdateUserPhoneNumberParams) (User, error)
	UpdateUserUsername(ctx context.Context, arg UpdateUserUsernameParams) (User, error)
}

var _ Querier = (*Queries)(nil)
