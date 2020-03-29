package todo

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Insert endpoint.Endpoint
	List   endpoint.Endpoint
}

func NewEndpoint(service Service) *Endpoints {
	return &Endpoints{
		Insert: Insert(service),
		List:   List(service),
	}
}

type (
	CreateTodoRequest struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	}
	CreateTodoResponse struct {
		Success bool `json:"success"`
	}
)

func Insert(service Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		requestModel := request.(CreateTodoRequest)
		err = service.Insert(requestModel.Title, requestModel.Text)
		responseModel := new(CreateTodoResponse)
		responseModel.Success = true
		return responseModel, err
	}
}

type (
	ListTodoRequest  struct{}
	ListTodoResponse struct {
		Success bool    `json:"success"`
		Data    []*TODO `json:"data"`
	}
)

func List(service Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		data, err := service.List()
		listTodoResponseModel := new(ListTodoResponse)
		listTodoResponseModel.Success = true
		listTodoResponseModel.Data = data
		return listTodoResponseModel, err
	}
}
