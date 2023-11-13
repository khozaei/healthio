package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Reception holds the schema definition for the Reception entity.
type Reception struct {
	ent.Schema
}

// Fields of the Reception.
func (Reception) Fields() []ent.Field {
	return []ent.Field{
		field.Time("reception_for"),
		field.Int("visit_duration").Positive(),
		field.String("insurance_code").Optional(),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Reception.
func (Reception) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("patient", Patient.Type).
			Ref("reception").Unique(),
		edge.To("visit", Visit.Type),
	}
}
