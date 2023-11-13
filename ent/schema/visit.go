package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Visit holds the schema definition for the Visit entity.
type Visit struct {
	ent.Schema
}

// Fields of the Visit.
func (Visit) Fields() []ent.Field {
	return []ent.Field{
		field.String("visit_price").Optional(),
		field.Time("visited_at").Optional(),
		field.String("payment_type").Optional(),
		field.Bool("is_paid").Optional(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Visit.
func (Visit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("reception", Reception.Type).
			Ref("visit").Unique(),
		edge.To("attachment", Attachment.Type),
	}
}
