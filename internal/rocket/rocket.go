package rocket

import "context"

// Rocket - should contain the definition of the Rocket
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// Store - Defines the methods that the Service will use to interact with the database
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(r Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service - Responsible for updating the rocket inventory
type Service struct {
	Store Store
}

// New - Creates a new instance of the Service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketByID - Retrieves a rocket by its ID
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// InsertRocket - Inserts a new rocket into the database
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// DeleteRocket - deletes a rocket from the inventory
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
