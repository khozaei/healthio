package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.String("father_name").Optional(),
		field.String("phone").Optional(),
		field.String("national_code").Unique().Optional(),
		field.String("identity_code").Unique().Optional(),
		field.Time("created_at").Immutable().Default(time.Now()),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attachment", Attachment.Type),
		edge.To("reception", Reception.Type),
	}
}
