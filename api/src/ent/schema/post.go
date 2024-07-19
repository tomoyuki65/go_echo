package schema

import (
    "time"

    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
    "entgo.io/ent/schema/edge"
)

type Post struct {
    ent.Schema
}

func (Post) Fields() []ent.Field {
    return []ent.Field{
        field.Int64("id"),
        field.Int64("user_id").
              Optional(),
        field.String("text").
              NotEmpty(),
        field.Time("created_at").
              SchemaType(map[string]string{
                  "mysql": "datetime",
              }).
              Default(time.Now).
              Immutable(),
        field.Time("updated_at").
              SchemaType(map[string]string{
                  "mysql": "datetime",
              }).
              Default(time.Now).
              UpdateDefault(time.Now),
        field.Time("deleted_at").
              SchemaType(map[string]string{
                  "mysql": "datetime",
              }).
              Nillable().
              Optional(),
    }
}

func (Post) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("user_id"),
        index.Fields("deleted_at"),
    }
}

func (Post) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("users", User.Type).
             Ref("posts").
             Field("user_id").
             Unique(),
    }
}
