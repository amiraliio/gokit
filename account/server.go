package account

import (
	httpTransport "github.com/go-kit/kit/transport/http"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints *Endpoints) http.Handler{
    req := mux.NewRouter()
    req.Use(commonMiddleware)

    req.Methods("POST").Path("/users").Handler(httpTransport.NewServer(
        endpoints.CreateUser,
        decodeUserReq,
        encodeResponse,
    ))

    req.Methods("GET").Path("/users/{id}").Handler(httpTransport.NewServer(
       endpoints.GetUser,
       decodeEmailReq,
       encodeResponse,
    ))

    return req
}

func commonMiddleware(next http.Handler) http.Handler{
    return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request){
        res.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(res, req)
    })
}