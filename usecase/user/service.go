package user

import (
	"context"
	"strings"
	"time"

	"github.com/chkilel/fiberent/entity"
)

//Service interface
type Service struct {
	repo Repository
}

//NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//GetUser Get a user by ID
func (s *Service) GetUser(ctx context.Context, id *entity.ID) (*entity.User, error) {
	return s.repo.Get(ctx, id)
}

//SearchUsers Search users
func (s *Service) SearchUsers(ctx context.Context, query string) ([]*entity.User, error) {
	return s.repo.Search(ctx, strings.ToLower(query))
}

//ListUsers List users
func (s *Service) ListUsers(ctx context.Context) ([]*entity.User, error) {
	return s.repo.List(ctx)
}

//CreateUser Create a user
func (s *Service) CreateUser(ctx context.Context, email, password, firstName, lastName string) (*entity.User, error) {

	_, err := s.repo.GetByEmail(ctx, email)
	if err == nil {
		return nil, entity.ErrEmailAlreadyRegistred
	}

	// Create a new user and generate encripted password
	e, err := entity.NewUser(email, password, firstName, lastName)
	if err != nil {
		return e, err
	}
	return s.repo.Create(ctx, e)
}

//UpdateUser Update a user
func (s *Service) UpdateUser(ctx context.Context, e *entity.User) (*entity.User, error) {

	password := strings.TrimSpace(e.Password)
	email := strings.TrimSpace(e.Email)
	var err error

	// User want to change password
	if password != "" {
		e.Password, err = e.GeneratePassword(password)
		if err != nil {
			return nil, err
		}
	}

	savedUser, err := s.repo.Get(ctx, &e.ID)
	if err != nil {
		return nil, err
	}

	// User want to change email
	if email != "" && email != savedUser.Email {

		// GetByEmail fails if no user found, or more than 1 user returned.
		_, err = s.repo.GetByEmail(ctx, email)
		if err == nil {
			return nil, entity.ErrEmailAlreadyRegistred
		}

		e.Email = email
	}

	e.UpdatedAt = time.Now()
	return s.repo.Update(ctx, e)
}

//DeleteUser Delete a user
func (s *Service) DeleteUser(ctx context.Context, id *entity.ID) error {
	u, err := s.GetUser(ctx, id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}
