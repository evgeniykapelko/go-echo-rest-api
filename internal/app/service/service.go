package service

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-echo-rest-api/internal/app/model"
	"log"
	"time"
)

type Service struct {
	DB *sql.DB
}

func New() *Service {
	connStr := "user=postgres password=postgres dbname=marketplace sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Service{DB: db}
}

func (s *Service) DaysLeft() int64 {
	d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	dur := time.Until(d)

	return int64(dur.Hours() / 24)
}

func (s *Service) GetAllCustomers() ([]*model.Customer, error) {
	query := `SELECT first_name, last_name, email, country FROM customers`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []*model.Customer
	for rows.Next() {
		var customer model.Customer
		if err := rows.Scan(&customer.FirstName, &customer.LastName, &customer.Email, &customer.Country); err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (s *Service) SaveCustomer(firstName, lastName, email, country, password string) error {
	query := `INSERT INTO customers (first_name, last_name, email, country, password) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.DB.Exec(query, firstName, lastName, email, country, password)
	return err
}

func (s *Service) DeleteCustomerByEmail(email string) error {
	query := `DELETE FROM customers WHERE email = $1`
	_, err := s.DB.Exec(query, email)
	return err
}

func (s *Service) GetCustomerByEmail(email string) ([]*model.Customer, error) {
	query := `SELECT first_name, last_name, email, country FROM customers WHERE email = $1`
	rows, err := s.DB.Query(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []*model.Customer
	for rows.Next() {
		var customer model.Customer
		if err := rows.Scan(&customer.FirstName, &customer.LastName, &customer.Email, &customer.Country); err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}
