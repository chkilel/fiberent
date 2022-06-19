package entity

import (
	"time"

	"github.com/chkilel/fiberent/ent"
)

type Pet struct {
	ent.Pet
}

func NewPet(name string, age int) *Pet {
	return &Pet{
		Pet: ent.Pet{
			Name:      name,
			Age:       age,
			CreatedAt: time.Now(),
		},
	}
}
