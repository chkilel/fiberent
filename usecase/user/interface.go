package user

import (
	"context"

	"github.com/chkilel/fiberent/entity"
)

//Reader interface
type Reader interface {
	Get(ctx context.Context, id *entity.ID) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Search(ctx context.Context, query string) ([]*entity.User, error)
	List(ctx context.Context) ([]*entity.User, error)
}

//Writer user writer
type Writer interface {
	Create(ctx context.Context, e *entity.User) (*entity.User, error)
	Update(ctx context.Context, e *entity.User) (*entity.User, error)
	Delete(ctx context.Context, id *entity.ID) error
	AddPets(ctx context.Context, userID *entity.ID, petIDs []*entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetUser(ctx context.Context, id *entity.ID) (*entity.User, error)
	SearchUsers(ctx context.Context, query string) ([]*entity.User, error)
	ListUsers(ctx context.Context) ([]*entity.User, error)
	CreateUser(ctx context.Context, email, password, firstName, lastName string) (*entity.User, error)
	UpdateUser(ctx context.Context, e *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id *entity.ID) error
	OwnPets(ctx context.Context, userID *entity.ID, petIDs []*entity.ID) error
}
