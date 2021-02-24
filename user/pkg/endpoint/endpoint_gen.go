// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "user/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	AddGenderEndpoint           endpoint.Endpoint
	GetGenderEndpoint           endpoint.Endpoint
	ListGendersEndpoint         endpoint.Endpoint
	RemoveGenderEndpoint        endpoint.Endpoint
	CreateUserEndpoint          endpoint.Endpoint
	GetUserByIDEndpoint         endpoint.Endpoint
	GetUserByEmailEndpoint      endpoint.Endpoint
	UpdateUserEmailEndpoint     endpoint.Endpoint
	UpdateUserPasswordEndpoint  endpoint.Endpoint
	UpdateUserInfoEndpoint      endpoint.Endpoint
	DeleteUserSoftEndpoint      endpoint.Endpoint
	RecoverUserEndpoint         endpoint.Endpoint
	DeleteUserPermanentEndpoint endpoint.Endpoint
	VerifyPasswordEndpoint      endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.UserService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		AddGenderEndpoint:           MakeAddGenderEndpoint(s),
		CreateUserEndpoint:          MakeCreateUserEndpoint(s),
		DeleteUserPermanentEndpoint: MakeDeleteUserPermanentEndpoint(s),
		DeleteUserSoftEndpoint:      MakeDeleteUserSoftEndpoint(s),
		GetGenderEndpoint:           MakeGetGenderEndpoint(s),
		GetUserByEmailEndpoint:      MakeGetUserByEmailEndpoint(s),
		GetUserByIDEndpoint:         MakeGetUserByIDEndpoint(s),
		ListGendersEndpoint:         MakeListGendersEndpoint(s),
		RecoverUserEndpoint:         MakeRecoverUserEndpoint(s),
		RemoveGenderEndpoint:        MakeRemoveGenderEndpoint(s),
		UpdateUserEmailEndpoint:     MakeUpdateUserEmailEndpoint(s),
		UpdateUserInfoEndpoint:      MakeUpdateUserInfoEndpoint(s),
		UpdateUserPasswordEndpoint:  MakeUpdateUserPasswordEndpoint(s),
		VerifyPasswordEndpoint:      MakeVerifyPasswordEndpoint(s),
	}
	for _, m := range mdw["AddGender"] {
		eps.AddGenderEndpoint = m(eps.AddGenderEndpoint)
	}
	for _, m := range mdw["GetGender"] {
		eps.GetGenderEndpoint = m(eps.GetGenderEndpoint)
	}
	for _, m := range mdw["ListGenders"] {
		eps.ListGendersEndpoint = m(eps.ListGendersEndpoint)
	}
	for _, m := range mdw["RemoveGender"] {
		eps.RemoveGenderEndpoint = m(eps.RemoveGenderEndpoint)
	}
	for _, m := range mdw["CreateUser"] {
		eps.CreateUserEndpoint = m(eps.CreateUserEndpoint)
	}
	for _, m := range mdw["GetUserByID"] {
		eps.GetUserByIDEndpoint = m(eps.GetUserByIDEndpoint)
	}
	for _, m := range mdw["GetUserByEmail"] {
		eps.GetUserByEmailEndpoint = m(eps.GetUserByEmailEndpoint)
	}
	for _, m := range mdw["UpdateUserEmail"] {
		eps.UpdateUserEmailEndpoint = m(eps.UpdateUserEmailEndpoint)
	}
	for _, m := range mdw["UpdateUserPassword"] {
		eps.UpdateUserPasswordEndpoint = m(eps.UpdateUserPasswordEndpoint)
	}
	for _, m := range mdw["UpdateUserInfo"] {
		eps.UpdateUserInfoEndpoint = m(eps.UpdateUserInfoEndpoint)
	}
	for _, m := range mdw["DeleteUserSoft"] {
		eps.DeleteUserSoftEndpoint = m(eps.DeleteUserSoftEndpoint)
	}
	for _, m := range mdw["RecoverUser"] {
		eps.RecoverUserEndpoint = m(eps.RecoverUserEndpoint)
	}
	for _, m := range mdw["DeleteUserPermanent"] {
		eps.DeleteUserPermanentEndpoint = m(eps.DeleteUserPermanentEndpoint)
	}
	for _, m := range mdw["VerifyPassword"] {
		eps.VerifyPasswordEndpoint = m(eps.VerifyPasswordEndpoint)
	}
	return eps
}
