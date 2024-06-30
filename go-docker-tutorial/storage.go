package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) (int, error)
	DeleteAccount(int) error
	UpdateAccount(*Account) error

	GetAccounts() ([]*Account, error)
	GetAccountById(int) (*Account, error)

	CreateUser(*User) error
	GetUserByEmail(string) (*User, error)
}

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) Init() (error, error) {

	return s.CreateAccountTable(), s.CreateUserTable()
}

func (s *PostgresStore) CreateUserTable() error {
	query := `create table if not exists users (
    id serial primary key,
	email varchar(50),
	password varchar(500),
    token varchar(500),
    account_id serial,
    created_at timestamp
    )`

	_, err := s.db.Exec(query)
	return err
}
func (s *PostgresStore) CreateAccountTable() error {
	query := `create table if not exists account (
    id serial primary key,
    first_name varchar(50),
    last_name varchar(50),
    number serial,
    balance serial,
    created_at timestamp
    )`

	_, err := s.db.Exec(query)
	return err
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "host = db port = 5432 user=postgres dbname=postgres password=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rsp, err := s.db.Query("Select * from account")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	accounts := []*Account{}
	for rsp.Next() {
		account, err := scanIntoAccount(rsp)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
func (s *PostgresStore) CreateUser(user *User) error {
	query := `insert into users
        (email, password, token,account_id,created_at)
        values ($1,$2, $3,$4,$5)`

	rsp, err := s.db.Exec(query, user.Email, user.EncryptedPassword, user.Token, user.AccountId, user.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", rsp)
	return nil
}
func (s *PostgresStore) CreateAccount(account *Account) (int, error) {
	query := `insert into account
        (first_name, last_name, number,balance,created_at)
        values ($1,$2, $3,$4,$5) returning id`

	var id int
	err := s.db.QueryRow(query, account.FirstName, account.LastName, account.Number, account.Balance, account.CreatedAt).Scan(&id)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}

	fmt.Printf("New account created with id: %d\n", id)
	return id, nil
}
func (s *PostgresStore) UpdateAccount(account *Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error {

	_, err := s.db.Query("Delete from account where id=$1", id)

	return err
}
func (s *PostgresStore) GetAccountById(id int) (*Account, error) {

	rsp, err := s.db.Query("select * from account where id=$1", id)
	if err != nil {
		return nil, err
	}
	for rsp.Next() {
		return scanIntoAccount(rsp)
	}
	return nil, fmt.Errorf("account %d not found", id)
}
func (s *PostgresStore) GetUserByEmail(email string) (*User, error) {

	rsp, err := s.db.Query("select * from users where email=$1", email)
	if err != nil {
		return nil, err
	}
	for rsp.Next() {
		return scanIntoUsers(rsp)
	}
	return nil, fmt.Errorf("user %s not found", email)
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {

	account := new(Account)
	err := rows.Scan(&account.Id, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt)
	return account, err
}
func scanIntoUsers(rows *sql.Rows) (*User, error) {

	user := new(User)
	err := rows.Scan(&user.Id, &user.Email, &user.EncryptedPassword, &user.Token, &user.AccountId, &user.CreatedAt)
	return user, err
}
