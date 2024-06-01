package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {

		}
	}

}

type ApiServer struct {
	listenAddr string
}

func NewApiServer(listenAddr string) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
	}
}

func (*ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", s.handleAccount)
}
func (*ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil

}
func (*ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil

}
func (*ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil

}
func (*ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil

}
func (*ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil

}
