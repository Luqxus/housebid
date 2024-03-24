package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/luquxSentinel/housebid/service"
	"github.com/luquxSentinel/housebid/types"
)

type APIFunc func(writer http.ResponseWriter, request *http.Request) error

type APIServer struct {
	mux           *http.ServeMux
	listenAddress string
	authService   service.AuthService
}

func NewAPIServer(listenAddress string, authService service.AuthService) *APIServer {
	return &APIServer{
		mux:           http.NewServeMux(),
		listenAddress: listenAddress,
		authService:   authService,
	}
}

func (api *APIServer) Run() error {

	// register user request
	http.HandleFunc("/signup", handlerFunc(api.registerUser))

	// login user request
	http.HandleFunc("/signin", handlerFunc(api.loginUser))

	//	start server and listen for requests
	return http.ListenAndServe(api.listenAddress, api.mux)
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
func (api *APIServer) registerUser(writer http.ResponseWriter, request *http.Request) error {
	// TODO: get & serialize request data
	reqBody := new(types.CreateUserData)
	err := decodeJSON(request.Body, reqBody)
	if err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return err
	}

	// TODO: call registerUser service
	err = api.authService.RegisterUser(request.Context(), reqBody)

	// TODO: error handling
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}

	// TODO: response
	return writeJSON(writer, http.StatusOK, map[string]string{"message": "user created successfully"})
}

//TODO: login user

func (api *APIServer) loginUser(writer http.ResponseWriter, request *http.Request) error {
	// get & decode request body
	reqBody := new(types.LoginData)
	err := decodeJSON(request.Body, reqBody)
	if err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return err
	}

	// call login service func
	user, err := api.authService.LoginUser(request.Context(), reqBody)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}

	// response
	return writeJSON(writer, http.StatusOK, map[string]any{"message": user})
}

//TODO: list a house

//TODO: bid on a house

//TODO: cancel a bid

//TODO:

func decodeJSON(r io.Reader, v any) error {
	// decode reader into v
	return json.NewDecoder(r).Decode(v)
}

func writeJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}
