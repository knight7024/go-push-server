package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/knight7024/go-push-server/common/util"
)

// DateTimeMixin holds the schema definition for the DateTimeMixin entity.
type DateTimeMixin struct {
	mixin.Schema
}

// Fields of the DateTimeMixin.
func (DateTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			GoType(util.Datetime("")).Immutable().Default(util.NowDatetime),
		field.Time("updated_at").
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
			}).
			GoType(util.Datetime("")).Immutable().Default(util.NowDatetime).UpdateDefault(util.NowDatetime),
	}
}
