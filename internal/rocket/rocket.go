package rocket

import (
	"context"
	"lessone-grpc-services/internal/domain"
)

type Store interface {
	GetById(id string) (domain.Rocket, error)
	Insert(r domain.Rocket) (domain.Rocket, error)
	Delete(id string) error
}

type servise struct {
	Store Store
}

func New(store Store) *servise {
	return &servise{
		Store: store,
	}
}

// Create Cocket
func (s *servise) GetById(ctx context.Context, id string) (*domain.Rocket, error) {
	rct, err := s.Store.GetById(id)
	if err != nil {
		return nil, err
	}
	return &rct, nil
}

// Insert Rocket
func (s *servise) Insert(ctx context.Context, r domain.Rocket) (*domain.Rocket, error) {
	rct, err := s.Store.Insert(r)
	if err != nil {
		return nil, err
	}
	return &rct, nil
}

// Delete Roket
func (s *servise) Delete(ctx context.Context, id string) error {
	err := s.Store.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
