package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//============================================================================//

// Model base model
type Model struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt uint64             `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt uint64             `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt uint64             `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

// BeforeCreate : handle before create
func (m *Model) BeforeCreate() {
	m.ID = primitive.NewObjectID()
	if m.CreatedAt == 0 {
		m.CreatedAt = uint64(time.Now().UnixMilli())
	}

	if m.UpdatedAt == 0 {
		m.UpdatedAt = uint64(time.Now().UnixMilli())
	}
}

// BeforeUpdate : handle before update
func (m *Model) BeforeUpdate() {
	m.UpdatedAt = uint64(time.Now().UnixMilli())
}

//============================================================================//
