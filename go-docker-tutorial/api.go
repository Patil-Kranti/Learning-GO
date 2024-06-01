package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string
}
type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusOK, ApiError{Error: err.Error()})

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

func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHttpHandleFunc(s.handleAccount))

	log.Println("Server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
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
