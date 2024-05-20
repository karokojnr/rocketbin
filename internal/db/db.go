package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/karokojnr/rocketbin/internal/rocket"
)

type Store struct {
	db *sqlx.DB
}

// New - creates a new Store
func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbTable := os.Getenv("DB_TABLE")
	dbSSLMode := os.Getenv("SSL_MODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbTable,
		dbPassword,
		dbSSLMode,
	)
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return Store{}, err
	}
	return Store{
		db: db,
	}, nil
}

// GetRocketByID - retrieves a rocket from the database and returns it or an error.
func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	var rkt rocket.Rocket
	row := s.db.QueryRow(
		`SELECT id FROM rockets WHERE id=(?)::uuid`,
		id,
	)
	err := row.Scan(&rkt)
	if err != nil {
		return rocket.Rocket{}, err
	}
	return rkt, nil
}

func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (s Store) DeleteRocket(id string) error {
	return nil
}
