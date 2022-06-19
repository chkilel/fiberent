package presenter

import "github.com/chkilel/fiberent/entity"

//User data
type Pet struct {
	ID   entity.ID `json:"id,omitempty"`
	Name string    `json:"first_name"`
	Age  int       `json:"age"`
}
