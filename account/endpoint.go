package account

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

//expose to public
type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s Service) *Endpoints {
	return &Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetsUserEndpoint(s),
	}
}

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}
)

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateUserRequest)
		id, err := s.CreateUser(ctx, req.Email, req.Password)
		return &CreateUserResponse{Ok: id}, err
	}
}

type (
	GetUserRequest struct {
		ID string `json:"id"`
	}
	GetUserResponse struct {
		Email string `json:"email"`
	}
)

func makeGetsUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserRequest)
		email, err := s.GetUser(ctx, req.ID)
		return &GetUserResponse{Email: email}, err
	}
}
