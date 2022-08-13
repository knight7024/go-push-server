package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"strings"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Projects.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("project_name").NotEmpty().MaxLen(128),
		field.String("project_id").NotEmpty().Unique(),
		field.Bytes("credentials").NotEmpty(),
		field.String("client_key").NotEmpty().DefaultFunc(func() string {
			clientKey, _ := uuid.NewRandom()
			return strings.ToUpper(strings.ReplaceAll(clientKey.String(), "-", ""))
		}).Immutable(),
		field.Int("user_id").Positive().StructTag(`json:"-"`),
	}
}

func (Project) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DateTimeMixin{},
		//SoftDeleteMixin{},
	}
}

// Edges of the Projects.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("projects").Field("user_id").Required().Unique(),
	}
}
