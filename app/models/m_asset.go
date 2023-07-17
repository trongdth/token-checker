package models

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
	ID          string         `json:"id" bson:"id"`
	Blockchain  *TwTBlockchain `json:"blockchain" bson:"blockchain"`
}
