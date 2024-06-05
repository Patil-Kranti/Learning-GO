package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error

	GetAccounts() ([]*Account, error)
	GetAccountById(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) Init() error {

	return s.CreateAccountTable()
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
	connStr := "host = localhost port = 5433 user=postgres dbname=postgres password=root sslmode=disable"

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
func (s *PostgresStore) CreateAccount(account *Account) error {
	query := `insert into account
        (first_name, last_name, number,balance,created_at)
        values ($1,$2, $3,$4,$5)`

	rsp, err := s.db.Exec(query, account.FirstName, account.LastName, account.Number, account.Balance, account.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", rsp)
	return nil
}
func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
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

func scanIntoAccount(rows *sql.Rows) (*Account, error) {

	account := new(Account)
	err := rows.Scan(&account.Id, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt)
	return account, err
}
