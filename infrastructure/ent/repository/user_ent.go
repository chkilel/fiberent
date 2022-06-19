package repository

import (
	"context"
	"strings"

	"github.com/chkilel/fiberent/ent"
	"github.com/chkilel/fiberent/ent/user"
	"github.com/chkilel/fiberent/entity"
	usecase "github.com/chkilel/fiberent/usecase/user"
)

//userRepoEnt Ent repo
type userRepoEnt struct {
	client *ent.Client
}

// NewUserRepoEnt is specific implementation of the interface
func NewUserRepoEnt(client *ent.Client) usecase.Repository {
	return &userRepoEnt{
		client: client,
	}
}

//Create a user
func (r *userRepoEnt) Create(ctx context.Context, e *entity.User) (*entity.User, error) {
	u, err := r.client.User.Create().
		SetFirstName(e.FirstName).
		SetLastName(e.LastName).
		SetEmail(e.Email).
		SetPassword(e.Password).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrCannotBeCreated
	}

	user := &entity.User{*u}

	return user, nil
}

//Get a user
func (r *userRepoEnt) Get(ctx context.Context, id *entity.ID) (*entity.User, error) {
	u, err := r.client.User.
		Query().
		Where(user.IDEQ(*id)).
		Only(ctx) // `Only` fails if no user found, or more than 1 user returned.

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.User{*u}, nil
}

//Get a user by email
func (r *userRepoEnt) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	u, err := r.client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx) // `Only` fails if no user found, or more than 1 user returned.

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.User{*u}, nil
}

//Update a user
func (r *userRepoEnt) Update(ctx context.Context, e *entity.User) (*entity.User, error) {

	// Prepare the update query
	query := r.client.User.
		UpdateOneID(e.ID)

	// Check if the user is updating his password
	if strings.TrimSpace(e.Password) != "" {
		query = query.SetPassword(e.Password)
	}

	u, err := query.
		SetFirstName(e.FirstName).
		SetLastName(e.LastName).
		SetEmail(e.Email).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return &entity.User{*u}, nil
}

//Delete a user
func (r *userRepoEnt) Delete(ctx context.Context, id *entity.ID) error {
	err := r.client.User.
		DeleteOneID(*id).
		Exec(ctx)

	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

//List users
func (r *userRepoEnt) List(ctx context.Context) ([]*entity.User, error) {
	entUsers, err := r.client.User.
		Query().
		All(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}

	users := make([]*entity.User, len(entUsers))
	for i, u := range entUsers {
		users[i] = &entity.User{*u}
	}

	return users, nil
}

//Search users
func (r *userRepoEnt) Search(ctx context.Context, query string) ([]*entity.User, error) {
	entUsers, err := r.client.User.
		Query().
		Where(user.FirstNameEQ(query)).
		All(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}

	users := make([]*entity.User, len(entUsers))
	for i, u := range entUsers {
		users[i] = &entity.User{*u}
	}

	return users, nil
}
