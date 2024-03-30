package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/luquxSentinel/housebid/middleware"
	"github.com/luquxSentinel/housebid/service"
	"github.com/luquxSentinel/housebid/types"
)

type APIFunc func(writer http.ResponseWriter, request *http.Request) error

type APIServer struct {
	mux           *http.ServeMux
	listenAddress string
	authService   service.AuthService
	houseService  service.HouseService
}

func NewAPIServer(listenAddress string, authService service.AuthService, houseService service.HouseService) *APIServer {
	return &APIServer{
		mux:           http.NewServeMux(),
		listenAddress: listenAddress,
		authService:   authService,
		houseService:  houseService,
	}
}

func (api *APIServer) Run() error {

	// register user request
	api.mux.HandleFunc("POST /signup", handlerFunc(api.registerUser))

	// login user request
	api.mux.HandleFunc("POST /signin", handlerFunc(api.loginUser))

	// list house request
	api.mux.HandleFunc("POST /list-house", middleware.AuthorizationMiddleware(handlerFunc(api.listHouse)))

	// get filtered houses request
	api.mux.HandleFunc("GET /houses/filter", middleware.AuthorizationMiddleware(handlerFunc(api.getHousesByFilter)))

	log.Printf("server listening on port : %s.", api.listenAddress)
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
		if err != nil {
			panic(err)
		}
	}
}

// TODO: register user
func (api *APIServer) registerUser(writer http.ResponseWriter, request *http.Request) error {
	// TODO: get & serialize request data
	reqBody := new(types.CreateUserData)
	err := decodeJSON(request.Body, reqBody)
	if err != nil {
		log.Println(err)
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return nil
	}

	// TODO: call registerUser service
	err = api.authService.RegisterUser(request.Context(), reqBody)

	// TODO: error handling
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return nil
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
		return nil
	}

	// call login service func
	user, token, err := api.authService.LoginUser(request.Context(), reqBody)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// set authorization header
	writer.Header().Set("authorization", token)

	// response
	return writeJSON(writer, http.StatusOK, map[string]any{"message": user})
}

// TODO: list a house
func (api *APIServer) listHouse(writer http.ResponseWriter, request *http.Request) error {
	// get uid from request context
	uid := request.Context().Value("uid").(string)
	if uid == "" {
		http.Error(writer, "user unauthorized user", http.StatusBadRequest)
		return nil
	}

	// get & decode request body
	reqBody := new(types.CreateHouseData)
	err := decodeJSON(request.Body, reqBody)
	if err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return nil
	}

	// call create house service func
	err = api.houseService.ListHouse(request.Context(), reqBody, uid)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// write response
	return writeJSON(writer, http.StatusOK, map[string]string{"message": "house listed successfully"})
}

func (api *APIServer) getHousesByFilter(writer http.ResponseWriter, request *http.Request) error {
	// get request data | filter
	filter := new(types.HouseQueryFilter)
	err := decodeJSON(request.Body, filter)
	if err != nil {
		http.Error(writer, "invalid filter body", http.StatusBadRequest)
		return nil
	}

	houses, err := api.houseService.GetHousesByFilter(request.Context(), filter)
	if err != nil {
		log.Printf("error occured while fectching houses with filter : %+v", filter)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return nil
	}

	err = writeJSON(writer, http.StatusOK, houses)
	return err
}

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
