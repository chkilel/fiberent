package pet

import (
	"context"

	"github.com/chkilel/fiberent/entity"
)

// Reader interface
type Reader interface {
	Get(ctx context.Context, id *entity.ID) (*entity.Pet, error)
	Search(ctx context.Context, query string) ([]*entity.Pet, error)
	List(ctx context.Context) ([]*entity.Pet, error)
}

// Writer user writer
type Writer interface {
	Create(ctx context.Context, e *entity.Pet) (*entity.Pet, error)
	Update(ctx context.Context, e *entity.Pet) (*entity.Pet, error)
	Delete(ctx context.Context, id *entity.ID) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	CreatePet(ctx context.Context, name string, age int) (*entity.Pet, error)
	GetPet(ctx context.Context, id *entity.ID) (*entity.Pet, error)
	UpdatePet(ctx context.Context, e *entity.Pet) (*entity.Pet, error)
	DeletePet(ctx context.Context, id *entity.ID) error
	SearchPets(ctx context.Context, query string) ([]*entity.Pet, error)
	ListPets(ctx context.Context) ([]*entity.Pet, error)
}
