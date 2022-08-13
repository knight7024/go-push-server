package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/knight7024/go-push-server/common/util"
)

// SoftDeleteMixin holds the schema definition for the SoftDeleteMixin entity.
type SoftDeleteMixin struct {
	mixin.Schema
}

// Fields of the SoftDeleteMixin.
func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").
			GoType(util.Datetime("")).Optional().Nillable(),
		field.String("deleted_by").Optional().Nillable().Sensitive().MaxLen(64),
	}
}
