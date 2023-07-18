package models

import (
	"time"
)

type TwTAsset struct {
	Model       `json:",inline" bson:",inline"`
	Name        string         `json:"name" bson:"name"`
	Type        string         `json:"type" bson:"type"`
	Symbol      string         `json:"symbol" bson:"symbol"`
	Decimals    uint32         `json:"decimals" bson:"decimals"`
	Website     string         `json:"website" bson:"website"`
	Description string         `json:"description" bson:"description"`
	Explorer    string         `json:"explorer" bson:"explorer"`
	Status      string         `json:"status" bson:"status"`
	Blockchain  *TwTBlockchain `json:"blockchain" bson:"blockchain"`
}

// BeforeCreate : handle before create
func (m *TwTAsset) BeforeCreate() {
	if m.CreatedAt == 0 {
		m.CreatedAt = uint64(time.Now().UnixMilli())
	}

	if m.UpdatedAt == 0 {
		m.UpdatedAt = uint64(time.Now().UnixMilli())
	}
}
