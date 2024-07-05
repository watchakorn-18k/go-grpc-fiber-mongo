package services

import context "context"

type UserService interface {
	GetUser(req *GetUserRequest) (*GetUserResponse, error)
}

type userServer struct {
	userServiceClient UserServiceClient
}

func NewUserServer(userServiceClient UserServiceClient) *userServer {
	return &userServer{userServiceClient}
}

func (sv *userServer) GetUser(req *GetUserRequest) (*GetUserResponse, error) {
	return sv.userServiceClient.GetUser(context.Background(), req)
}
