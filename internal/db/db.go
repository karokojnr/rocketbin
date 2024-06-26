package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/karokojnr/rocketbin/internal/rocket"
	uuid "github.com/satori/go.uuid"
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
		`SELECT id, type, name FROM rockets WHERE id=$1;`,
		id,
	)
	err := row.Scan(&rkt.ID, &rkt.Name, &rkt.Type)
	if err != nil {
		return rocket.Rocket{}, err
	}
	return rkt, nil
}

// InsertRocket - inserts a rocket into the rockets table
func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	_, err := s.db.NamedQuery(
		`INSERT INTO rockets
		(id, name, type)
		VALUES (:id, :name, :type)`,
		rkt,
	)
	if err != nil {
		return rocket.Rocket{}, errors.New("failed to insert into database")
	}
	return rocket.Rocket{
		ID:   rkt.ID,
		Type: rkt.Type,
		Name: rkt.Name,
	}, nil
}

// DeleteRocket - attempts to delete a rocket from the database return err if error
func (s Store) DeleteRocket(id string) error {
	uid, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(
		`DELETE FROM rockets where id = $1`,
		uid,
	)
	if err != nil {
		return err
	}
	return nil
}
