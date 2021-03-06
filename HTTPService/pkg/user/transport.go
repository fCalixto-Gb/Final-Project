package user

import (
	"context"
	"encoding/json"
	"net/http"

	ent "github.com/fCalixto-Gb/Final-Project/HTTPService/pkg/entities"
	erro "github.com/fCalixto-Gb/Final-Project/HTTPService/pkg/errors"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

var (
	GetUserRoute      = "/User/{Id}"
	CreateUserRoute   = "/User"
	DeleteUserRoute   = "/User/{Id}"
	UpdateUserRoute   = "/User"
	AuthenticateRoute = "/Authenticate"
)

func NewMuxApi(endpoints Endpoints, logger log.Logger) http.Handler {
	// options to support error control
	options := []kitHttp.ServerOption{
		kitHttp.ServerErrorLogger(logger),
		kitHttp.ServerErrorEncoder(encodeErrorResponse), // encode error to the http.ResponseWriter whenever one is found processing a request
	}
	r := mux.NewRouter()
	// GetUser method handler
	r.Methods("GET").Path(GetUserRoute).Handler(
		kitHttp.NewServer(
			endpoints.GetUser,
			decodeGetUser,
			encodeStatusOKResp,
			options...,
		),
	)
	// CreateUser method handler
	r.Methods("POST").Path(CreateUserRoute).Handler(
		kitHttp.NewServer(
			endpoints.CreateUser,
			decodeCreateUser,
			encodeCreateResp,
			options...,
		),
	)
	// CreateUser method handler
	r.Methods("DELETE").Path(DeleteUserRoute).Handler(
		kitHttp.NewServer(
			endpoints.DeleteUser,
			decodeDeleteUser,
			encodeStatusOKResp,
			options...,
		),
	)
	// updateUser method handler
	r.Methods("PUT").Path(UpdateUserRoute).Handler(
		kitHttp.NewServer(
			endpoints.UpdateUser,
			decodeUpdateUser,
			encodeStatusOKResp,
			options...,
		),
	)
	// Authenticate method handler
	r.Methods("POST").Path(AuthenticateRoute).Handler(
		kitHttp.NewServer(
			endpoints.Authenticate,
			decodeAuthenticateUser,
			encodeStatusOKResp,
			options...,
		),
	)
	return r
}

// Encodes to json to give a http response
func encodeStatusOKResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

// Encodes to json to give a 201 created response
func encodeCreateResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(response)
}

// decode request entering the http server
func decodeGetUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ent.GetUserReq
	// Getting passed id from path params
	Id := mux.Vars(r)["Id"]
	req.Id = Id
	return req, nil
}

// decode request entering the http server
func decodeDeleteUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ent.DeleteUserReq
	// Getting passed id from path params
	Id := mux.Vars(r)["Id"]
	req.Id = Id
	return req, nil
}

// decode request entering the http server
func decodeCreateUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ent.CreateUserReq
	// translating json to my CreateUserReq struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// decode request entering the http server
func decodeUpdateUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ent.UpdateUserReq
	// translating json to my CreateUserReq struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// decode request entering the http server
func decodeAuthenticateUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ent.AuthenticateReq
	// translating json to my CreateUserReq struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// encode my custom errors to json response
func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(erro.ErrToHTTPStatus(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
