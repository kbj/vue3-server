package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").NotEmpty().MaxLen(50).Comment("用户名"),
		field.Text("realname").NotEmpty().MaxLen(50).Comment("真实名称"),
		field.Text("cellphone").MinLen(11).MaxLen(11).Comment("手机号"),
		field.Int8("enable").Max(9).Default(1).Comment("是否启用，1启用0禁用"),
		field.Time("create_at").StructTag(`json:"createAt,omitempty"`).Comment("创建时间").Default(func() time.Time {
			return time.Now()
		}).Immutable(),
		field.Time("update_at").StructTag(`json:"updateAt,omitempty"`).Comment("修改时间").Default(func() time.Time {
			return time.Now()
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
