package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddr string
	store      Storage
}

func NewApiServer(listenAddr string, store Storage) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHttpHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", withJWTAuth(makeHttpHandleFunc(s.handleGetAccountById), s.store))
	router.HandleFunc("/transfer", makeHttpHandleFunc(s.handleTransfer))

	log.Println("Server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}
func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "PUT" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("Method not allowed: %s", r.Method)

}
func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountRequest := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccountRequest); err != nil {
		return err
	}
	account := NewAccount(createAccountRequest.FirstName, createAccountRequest.LastName)

	if err := s.store.CreateAccount(account); err != nil {
		return err
	}
	tokenString, err := createJwt(account)
	if err != nil {
		return err
	}
	fmt.Println(tokenString)
	return WriteJson(w, http.StatusOK, account)

}
func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {

	account, err := s.store.GetAccounts()
	if err != nil {
		log.Fatal(err)
	}

	return WriteJson(w, http.StatusOK, account)

}
func (s *ApiServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {

		id, err := getIdIntFromString(r)
		if err != nil {
			return err
		}
		account, err := s.store.GetAccountById(id)
		if err != nil {
			return err
		}
		return WriteJson(w, http.StatusOK, account)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("Method not allowed: %s", r.Method)
}
func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {

	id, err := getIdIntFromString(r)

	if err != nil {
		return err
	}

	if err := s.store.DeleteAccount(id); err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, map[string]int{"deleted": id})

}
func (s *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	transferReq := new(TransferRequest)
	if err := json.NewDecoder(r.Body).Decode(transferReq); err != nil {
		return err
	}

	defer r.Body.Close()

	return WriteJson(w, http.StatusOK, transferReq)

}
func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string `json:"error"`
}

func createJwt(account *Account) (string, error) {

	claims := &jwt.MapClaims{
		"expiresAt":     jwt.NewNumericDate(time.Unix(1516239022, 0)),
		"accountNumber": account.Number,
	}
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
func permissionDenied(w http.ResponseWriter) {

	WriteJson(w, http.StatusForbidden, ApiError{Error: "Permission Denied"})
}
func withJWTAuth(handlerFunc http.HandlerFunc, s Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Calling the JWT middleware")
		tokenString := r.Header.Get("x-jwt-token")
		token, error := validateJWT(tokenString)

		if error != nil {
			permissionDenied(w)
			return
		}

		if !token.Valid {
			permissionDenied(w)
			return
		}
		userId, err := getIdIntFromString(r)
		if err != nil {
			permissionDenied(w)
			return
		}
		account, error := s.GetAccountById(userId)
		if error != nil {
			permissionDenied(w)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(int64(claims["accountNumber"].(float64)))
		fmt.Println(account.Number)

		if account.Number != int64(claims["accountNumber"].(float64)) {
			permissionDenied(w)
			return
		}
		handlerFunc(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusOK, ApiError{Error: err.Error()})

		}
	}

}

func getIdIntFromString(r *http.Request) (int, error) {

	idstr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Fatal(err)
		return id, fmt.Errorf("Id %s was invalid", idstr)
	}
	return id, nil
}
