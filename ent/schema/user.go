package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty().MinLen(4).MaxLen(64).Unique(),
		field.String("password").NotEmpty().Sensitive(),
		field.Bool("is_approved").Immutable().StructTag(`json:"-"`).Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("projects", Project.Type),
	}
}
