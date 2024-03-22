package main

import (
	"context"
	"log"
	"net/http"
	"time"
)



type APIFunc func(writer http.ResponseWriter, request *http.Request) error

type APIServer struct {
	mux *http.ServeMux
	listenAddress string
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		mux: http.NewServeMux(),
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
		ctx, cancel := context.WithTimeout(context.Background(), 30 *time.Second)

//		cancel context
		defer cancel()

//		handler helper
		err := fn(writer, request.WithContext(ctx))

//		on error panic
		log.Panic(err)
	}
}