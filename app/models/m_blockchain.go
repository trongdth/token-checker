package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TwTBlockchain struct {
	Model       `json:",inline" bson:",inline"`
	Name        string `json:"name" bson:"name"`
	Website     string `json:"website" bson:"website"`
	Description string `json:"description" bson:"description"`
	Explorer    string `json:"explorer" bson:"explorer"`
	Research    string `json:"research" bson:"research"`
	Symbol      string `json:"symbol" bson:"symbol"`
	Type        string `json:"type" bson:"type"`
	Status      string `json:"status" bson:"status"`
	Decimals    uint32 `json:"decimals" bson:"decimals"`
}

// BeforeCreate : handle before create
func (m *TwTBlockchain) BeforeCreate() {
	m.ID = primitive.NewObjectID().String()
	if m.CreatedAt == 0 {
		m.CreatedAt = uint64(time.Now().UnixMilli())
	}

	if m.UpdatedAt == 0 {
		m.UpdatedAt = uint64(time.Now().UnixMilli())
	}
}
