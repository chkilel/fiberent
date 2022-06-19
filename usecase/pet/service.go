package pet

import (
	"context"

	"github.com/chkilel/fiberent/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreatePet Create a pet
func (s *Service) CreatePet(ctx context.Context, name string, age int) (*entity.Pet, error) {
	p := entity.NewPet(name, age)
	return s.repo.Create(ctx, p)
}

// GetPet Get a pet by ID
func (s *Service) GetPet(ctx context.Context, id *entity.ID) (*entity.Pet, error) {
	return s.repo.Get(ctx, id)
}

// UpdatePet Update a pet
func (s *Service) UpdatePet(ctx context.Context, p *entity.Pet) (*entity.Pet, error) {
	_, err := s.repo.Get(ctx, &p.ID)
	if err != nil {
		return nil, entity.ErrNotFound
	}

	return s.repo.Update(ctx, p)
}

// DeletePet Delete a pet
func (s *Service) DeletePet(ctx context.Context, id *entity.ID) error {

	p, err := s.repo.Get(ctx, id)
	if p == nil || err != nil {
		return entity.ErrNotFound
	}
	return s.repo.Delete(ctx, id)
}

// ListPets List pets
func (s *Service) ListPets(ctx context.Context) ([]*entity.Pet, error) {
	return s.repo.List(ctx)
}

// SearchPets Search pets
func (s *Service) SearchPets(ctx context.Context, query string) ([]*entity.Pet, error) {
	return s.repo.Search(ctx, query)
}
