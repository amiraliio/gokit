package todo

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewTransport(ctx context.Context, endpoints *Endpoints) http.Handler {

	router := mux.NewRouter()

	router.Methods("GET").Path("/todo").Handler(httptransport.NewServer(
		endpoints.Insert,
		insertRequest,
		insertResponse,
	))

	router.Methods("POST").Path("/todo").Handler(httptransport.NewServer(
		endpoints.List,
		listRequest,
		listResponse,
	))

	return router
}

func insertRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var createTodoRequest CreateTodoRequest
	if err := json.NewDecoder(request.Body).Decode(&createTodoRequest); err != nil {
		return nil, err
	}
	return createTodoRequest, nil
}

func insertResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func listRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func listResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
