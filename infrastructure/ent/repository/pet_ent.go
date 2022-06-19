package repository

import (
	"context"

	"github.com/chkilel/fiberent/ent"
	"github.com/chkilel/fiberent/ent/pet"
	"github.com/chkilel/fiberent/entity"
	usecase "github.com/chkilel/fiberent/usecase/pet"
)

//petRepoEnt Ent repo
type petRepoEnt struct {
	client *ent.Client
}

// NewPetRepoEnt is specific implementation of the interface
func NewPetRepoEnt(client *ent.Client) usecase.Repository {
	return &petRepoEnt{
		client: client,
	}
}

//Create a pet
func (r *petRepoEnt) Create(ctx context.Context, p *entity.Pet) (*entity.Pet, error) {
	EntPet, err := r.client.Pet.Create().
		SetName(p.Name).
		SetAge(p.Age).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrCannotBeCreated
	}

	pet := &entity.Pet{*EntPet}

	return pet, nil
}

//Get a pet
func (r *petRepoEnt) Get(ctx context.Context, id *entity.ID) (*entity.Pet, error) {
	entPet, err := r.client.Pet.
		Query().
		Where(pet.IDEQ(*id)).
		Only(ctx) // `Only` fails if no pet found, or more than 1 pet returned.

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.Pet{*entPet}, nil
}

//Update a pet
func (r *petRepoEnt) Update(ctx context.Context, p *entity.Pet) (*entity.Pet, error) {

	// Prepare the update query
	entPet, err := r.client.Pet.
		UpdateOneID(p.ID).
		SetName(p.Name).
		SetAge(p.Age).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return &entity.Pet{*entPet}, nil
}

//Delete a pet
func (r *petRepoEnt) Delete(ctx context.Context, id *entity.ID) error {
	err := r.client.Pet.
		DeleteOneID(*id).
		Exec(ctx)

	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

//List pets
func (r *petRepoEnt) List(ctx context.Context) ([]*entity.Pet, error) {
	entPets, err := r.client.Pet.
		Query().
		All(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}

	pets := make([]*entity.Pet, len(entPets))
	for i, p := range entPets {
		pets[i] = &entity.Pet{*p}
	}

	return pets, nil
}

//Search pets
func (r *petRepoEnt) Search(ctx context.Context, name string) ([]*entity.Pet, error) {
	entPets, err := r.client.Pet.
		Query().
		Where(pet.NameEQ(name)).
		All(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}

	pets := make([]*entity.Pet, len(entPets))
	for i, p := range entPets {
		pets[i] = &entity.Pet{*p}
	}

	return pets, nil
}
