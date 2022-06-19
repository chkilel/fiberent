package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// ent framework has validator build in, the code below indicate that we define a field
		// with type of string named "first_name" and we validate it as required/ not empty
		// Giving the field `last_name` of type string to have default value of unkown if no value is supplied

		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
		field.String("first_name").NotEmpty(),
		field.String("last_name").Default("unknown"),
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").UpdateDefault(time.Now()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type),
	}
}
