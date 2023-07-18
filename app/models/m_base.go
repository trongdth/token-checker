package models

import (
	"time"
)

//============================================================================//

// Model base model
type Model struct {
	ID        string `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt uint64 `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt uint64 `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt uint64 `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

// BeforeUpdate : handle before update
func (m *Model) BeforeUpdate() {
	m.UpdatedAt = uint64(time.Now().UnixMilli())
}

//============================================================================//
