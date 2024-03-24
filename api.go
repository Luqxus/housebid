package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/luquxSentinel/housebid/types"
)

type APIFunc func(writer http.ResponseWriter, request *http.Request) error

type APIServer struct {
	mux           *http.ServeMux
	listenAddress string
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		mux:           http.NewServeMux(),
		listenAddress: listenAddress,
	}
}

func (s *APIServer) Run() error {

	//	start server and listen for requests
	return http.ListenAndServe(s.listenAddress, s.mux)
}

func handlerFunc(fn APIFunc) http.HandlerFunc {
	//	handle incoming request
	return func(writer http.ResponseWriter, request *http.Request) {

		//		timeout  context
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

		//		cancel context
		defer cancel()

		//		handler helper
		err := fn(writer, request.WithContext(ctx))

		//		on error panic
		log.Panic(err)
	}
}

// TODO: register user
func (api *APIServer) registerUser(writer http.ResponseWriter, request *http.Request) {
	// TODO: get & serialize request data
	reqBody := new(types.CreateUserData)
	err := decodeJSON(request.Body, reqBody)
	if err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return
	}

	// TODO: call registerUser service

	// TODO: error handling

	// TODO: response
}

//TODO: login user

//TODO: list a house

//TODO: bid on a house

//TODO: cancel a bid

//TODO:

func decodeJSON(r io.Reader, v any) error {
	// decode reader into v
	return json.NewDecoder(r).Decode(v)
}
