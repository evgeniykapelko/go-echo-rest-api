package service

import (
	"database/sql"
	_ "github.com/lib/pq"
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

func (s *Service) SaveCustomer(firstName, lastName, email, country, password string) error {
	query := `INSERT INTO customers (first_name, last_name, email, country, password) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.DB.Exec(query, firstName, lastName, email, country, password)
	return err
}
