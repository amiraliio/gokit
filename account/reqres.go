package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func encodeResponse(ctx context.Context, res http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(res).Encode(response)
}

func decodeUserReq(ctx context.Context, req *http.Request) (interface{}, error) {
	var userReq CreateUserRequest
	if err := json.NewDecoder(req.Body).Decode(&userReq); err != nil {
		return nil, err
	}
	return userReq, nil
}

func decodeEmailReq(ctx context.Context, req *http.Request) (interface{}, error) {
	var userReq GetUserRequest
	vars := mux.Vars(req)
	userReq.ID = vars["id"]
	return userReq, nil
}
