package presenter

import "github.com/chkilel/fiberent/entity"

//User data
type User struct {
	ID        entity.ID `json:"id,omitempty"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}
