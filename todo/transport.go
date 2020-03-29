package todo

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewTransport(_ context.Context, endpoints *Endpoints) http.Handler {

	router := mux.NewRouter()
	router.Use(commonMiddleware)

	router.Methods("GET").Path("/todo").Handler(httptransport.NewServer(
		endpoints.List,
		listRequest,
		listResponse,
	))

	router.Methods("POST").Path("/todo").Handler(httptransport.NewServer(
		endpoints.Insert,
		insertRequest,
		insertResponse,
	))

	return router
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
	return nil, nil
}

func listResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
