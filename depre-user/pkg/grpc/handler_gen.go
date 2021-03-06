// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "user/pkg/endpoint"
	pb "user/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	addGender           grpc.Handler
	getGender           grpc.Handler
	listGenders         grpc.Handler
	removeGender        grpc.Handler
	createUser          grpc.Handler
	getUserByID         grpc.Handler
	getUserByEmail      grpc.Handler
	updateUserEmail     grpc.Handler
	updateUserPassword  grpc.Handler
	updateUserInfo      grpc.Handler
	deleteUserSoft      grpc.Handler
	recoverUser         grpc.Handler
	deleteUserPermanent grpc.Handler
	verifyPassword      grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.UserServer {
	return &grpcServer{
		addGender:           makeAddGenderHandler(endpoints, options["AddGender"]),
		createUser:          makeCreateUserHandler(endpoints, options["CreateUser"]),
		deleteUserPermanent: makeDeleteUserPermanentHandler(endpoints, options["DeleteUserPermanent"]),
		deleteUserSoft:      makeDeleteUserSoftHandler(endpoints, options["DeleteUserSoft"]),
		getGender:           makeGetGenderHandler(endpoints, options["GetGender"]),
		getUserByEmail:      makeGetUserByEmailHandler(endpoints, options["GetUserByEmail"]),
		getUserByID:         makeGetUserByIDHandler(endpoints, options["GetUserByID"]),
		listGenders:         makeListGendersHandler(endpoints, options["ListGenders"]),
		recoverUser:         makeRecoverUserHandler(endpoints, options["RecoverUser"]),
		removeGender:        makeRemoveGenderHandler(endpoints, options["RemoveGender"]),
		updateUserEmail:     makeUpdateUserEmailHandler(endpoints, options["UpdateUserEmail"]),
		updateUserInfo:      makeUpdateUserInfoHandler(endpoints, options["UpdateUserInfo"]),
		updateUserPassword:  makeUpdateUserPasswordHandler(endpoints, options["UpdateUserPassword"]),
		verifyPassword:      makeVerifyPasswordHandler(endpoints, options["VerifyPassword"]),
	}
}
