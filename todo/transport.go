package todo

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewTransport(_ context.Context, endpoints *Endpoints) {

	http.Handle("/todo", httptransport.NewServer(
		endpoints.Insert,
		insertRequest,
		insertResponse,
	))

	http.Handle("/Todo", httptransport.NewServer(
		endpoints.List,
		listRequest,
		listResponse,
	))

	log.Fatal(http.ListenAndServe(":8976", nil))

}

func insertRequest(_ context.Context, request *http.Request) (interface{}, error) {
	var createTodoRequest CreateTodoRequest
	if err := json.NewDecoder(request.Body).Decode(&createTodoRequest); err != nil {
		return nil, err
	}
	return createTodoRequest, nil
}

func insertResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func listRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var listRequest ListTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&listRequest); err != nil {
		return nil, err
	}
	return listRequest, nil
}

func listResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
